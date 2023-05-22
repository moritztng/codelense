package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/messaging"
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
	conf := messaging.ReadConfig("kafka.properties")
	producer, _ := kafka.NewProducer(&conf)
	client := graphql.NewClient("https://api.github.com/graphql",
		&http.Client{Transport: &authedTransport{wrapped: http.DefaultTransport}})
	response, _ := getRepositories(context.Background(), client)
	edges := response.Search.Edges
	for _, edge := range edges {
		repository := edge.Node.(*getRepositoriesSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository)
		repositoryJson, _ := json.Marshal(messaging.Repository{Owner: repository.Owner.GetLogin(), Name: repository.Name, Stars: uint(repository.StargazerCount)})
		topic := "github_load_repositories"
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          repositoryJson,
		}, nil)
	}
	producer.Flush(15 * 1000)
	producer.Close()
}

//go:generate go run github.com/Khan/genqlient
