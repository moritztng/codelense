package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	minCreated := "1970-01-01"
	hasEdges := true
	total := 0
	for hasEdges {
		hasNextPage := true
		nextPageCursor := ""
		query := fmt.Sprintf("location:Germany type:org sort:joined-asc created:>%s", minCreated)
		for hasNextPage {
			response, _ := getOrganizations(context.Background(), client, nextPageCursor, query)
			edges := response.Search.Edges
			if len(edges) == 0 {
				hasEdges = false
				break
			}
			for _, edge := range edges {
				organization := edge.Node.(*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization)
				organizationJson, _ := json.Marshal(messaging.Repository{Key: organization.DatabaseId, Login: organization.Login, Name: organization.Name, CreatedAt: organization.CreatedAt})
				topic := "github_load_organizations"
				producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          organizationJson,
				}, nil)
				minCreated = organization.CreatedAt.Format("2006-01-02")
				total += 1
			}
			fmt.Println(minCreated)
			fmt.Println(total)
			hasNextPage = response.Search.PageInfo.HasNextPage
			nextPageCursor = response.Search.PageInfo.EndCursor
		}
	}
	producer.Flush(15 * 1000)
	producer.Close()
}

//go:generate go run github.com/Khan/genqlient
