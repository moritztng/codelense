package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/model"
	"go.uber.org/zap"
)

type authedTransport struct {
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	key := os.Getenv("GITHUB_TOKEN")
	req.Header.Set("Authorization", "bearer "+key)
	return t.wrapped.RoundTrip(req)
}

func main() {
	baseLogger, _ := zap.NewProduction()
	logger := baseLogger.Sugar()
	defer logger.Sync()
	logger.Info("start")
	kafkaProducer, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))})
	graphqlClient := graphql.NewClient("https://api.github.com/graphql",
		&http.Client{Transport: &authedTransport{wrapped: http.DefaultTransport}})
	kafkaConfig := kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
		"group.id":          "github_load_organization",
		"auto.offset.reset": "earliest",
	}
	kafkaConsumer, _ := kafka.NewConsumer(&kafkaConfig)
	defer kafkaConsumer.Close()
	kafkaConsumer.SubscribeTopics([]string{"github_organization_events"}, nil)
	for {
		message, err := kafkaConsumer.ReadMessage(time.Second)
		if err != nil {
			continue
		}
		var organizationEvent model.OrganizationEvent
		json.Unmarshal(message.Value, &organizationEvent)
		response, _ := getOrganization(context.Background(), graphqlClient, organizationEvent.Login)
		organization := response.GetOrganization()
		organizationJson, _ := json.Marshal(model.Organization{GithubID: uint(organization.DatabaseId), Login: organization.Login, Name: organization.Name, Location: organization.GetLocation(), GithubCreatedAt: organization.CreatedAt, Email: organization.Email, AvatarUrl: organization.AvatarUrl, Description: organization.Description, TwitterUsername: organization.TwitterUsername, GithubUpdatedAt: organization.UpdatedAt, WebsiteUrl: organization.WebsiteUrl, Url: organization.Url})
		topic := "github_load_organization"
		kafkaProducer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          organizationJson,
		}, nil)
		logger.Infow("loaded organization", "login", organization.Login)
	}
	logger.Info("stop")
	kafkaProducer.Close()
	kafkaProducer.Flush(15 * 1000)
}

//go:generate go run github.com/Khan/genqlient
