package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/go-github/v52/github"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/model"
	"golang.org/x/oauth2"
	"gorm.io/datatypes"
)

func main() {
	kafkaProducer, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	topic := "github_load_events"
	ctx := context.Background()
	token := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tokenClient := oauth2.NewClient(ctx, token)
	githubClient := github.NewClient(tokenClient)
	githubOptions := &github.ListOptions{PerPage: 100}
	lastId := ""
	duration, _ := time.ParseDuration(os.Getenv("INTERVAL"))
	ticker := time.NewTicker(duration)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			events, _, err := githubClient.Activity.ListEvents(context.Background(), githubOptions)
			if err != nil {
				fmt.Println(err)
				break
			}
			nEvents := 0
			for _, event := range events {
				if event.GetID() == lastId {
					fmt.Println("already loaded")
					break
				}
				id, _ := strconv.Atoi(event.GetID())
				eventJson, _ := json.Marshal(model.Event{GithubID: uint(id), Type: event.GetType(), ActorID: uint(event.GetActor().GetID()), OrgID: uint(event.GetOrg().GetID()), RepositoryID: uint(event.GetRepo().GetID()), Payload: datatypes.JSON(event.GetRawPayload()), GithubCreatedAt: event.GetCreatedAt().Time})
				kafkaProducer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          eventJson,
				}, nil)
				nEvents++
			}
			fmt.Println(nEvents)
			lastId = events[0].GetID()
		case <-quit:
			ticker.Stop()
			return
		}
	}
	kafkaProducer.Flush(15 * 1000)
	kafkaProducer.Close()
}
