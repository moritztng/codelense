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
	topic := "github"
	conf := messaging.ReadConfig("kafka.properties")
	p, _ := kafka.NewProducer(&conf)
	ctx := context.Background()
	client := graphql.NewClient("https://api.github.com/graphql",
		&http.Client{Transport: &authedTransport{wrapped: http.DefaultTransport}})
	resp, _ := getRepositories(ctx, client)
	repos := resp.Search.Edges
	for i := 0; i < len(repos); i++ {
		type Repo struct {
			Name  string
			Stars uint
		}
		repo := repos[i].Node.(*getRepositoriesSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository)
		repoJson, _ := json.Marshal(Repo{repo.Name, uint(repo.StargazerCount)})
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          repoJson,
		}, nil)
	}
	p.Flush(15 * 1000)
	p.Close()
}

//go:generate go run github.com/Khan/genqlient
