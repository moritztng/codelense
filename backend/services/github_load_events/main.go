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
	"github.com/moritztng/codelense/backend/util"
	"golang.org/x/oauth2"
)

func main() {
	conf := util.ReadConfig("kafka.properties")
	producer, _ := kafka.NewProducer(&conf)
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
				key, _ := strconv.Atoi(event.GetID())
				eventJson, _ := json.Marshal(util.Event{Key: key, Type: event.GetType(), ActorId: int(event.GetActor().GetID()), OrgId: int(event.GetOrg().GetID()), RepositoryId: int(event.GetRepo().GetID()), Payload: event.GetRawPayload(), CreatedAt: event.GetCreatedAt().Time})
				producer.Produce(&kafka.Message{
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
	producer.Flush(15 * 1000)
	producer.Close()
}

//go:generate go run github.com/Khan/genqlient
