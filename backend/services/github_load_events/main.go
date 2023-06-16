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
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

func main() {
	baseLogger, _ := zap.NewProduction()
	logger := baseLogger.Sugar()
	defer logger.Sync()
	logger.Info("start")
	kafkaProducer, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))})
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
			events, _, _ := githubClient.Activity.ListEvents(context.Background(), githubOptions)
			nEvents := 0
			for _, event := range events {
				if event.GetID() == lastId {
					break
				}
				id, _ := strconv.Atoi(event.GetID())
				eventJson, _ := json.Marshal(model.Event{GithubId: uint(id), Type: event.GetType(), ActorId: uint(event.GetActor().GetID()), ActorLogin: event.GetActor().GetLogin(), ActorUrl: event.GetActor().GetURL(), ActorAvatarUrl: event.GetActor().GetAvatarURL(), RepositoryId: uint(event.GetRepo().GetID()), RepositoryName: event.GetRepo().GetName(), RepositoryUrl: event.GetRepo().GetURL(), Payload: event.GetRawPayload(), Public: event.GetPublic(), GithubCreatedAt: event.GetCreatedAt().Time, OrgId: uint(event.GetOrg().GetID()), OrgLogin: event.GetOrg().GetLogin(), OrgUrl: event.GetOrg().GetURL(), OrgAvatarUrl: event.GetOrg().GetAvatarURL()})
				kafkaProducer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          eventJson,
				}, nil)
				nEvents++
			}
			logger.Infow("events loaded", "number", nEvents)
			lastId = events[0].GetID()
		case <-quit:
			ticker.Stop()
			return
		}
	}
	logger.Info("stop")
	kafkaProducer.Flush(15 * 1000)
	kafkaProducer.Close()
}
