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
	createdAfterMin := true
	total := 0
	for createdAfterMin {
		createdAfterMin = false
		hasNextPage := true
		nextPageCursor := ""
		query := fmt.Sprintf("location:Germany type:org sort:joined-asc created:>=%s", minCreated)
		for hasNextPage {
			response, _ := getOrganizations(context.Background(), client, nextPageCursor, query)
			edges := response.Search.Edges
			for _, edge := range edges {
				organization := edge.Node.(*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization)
				organizationJson, _ := json.Marshal(messaging.Organization{Key: organization.DatabaseId, Login: organization.Login, Name: organization.Name, CreatedAt: organization.CreatedAt})
				topic := "github_load_organizations"
				producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          organizationJson,
				}, nil)
				created := organization.CreatedAt.Format("2006-01-02")
				if created != minCreated {
					minCreated = created
					createdAfterMin = true
					fmt.Println(minCreated)
				}
				total += 1
			}
			fmt.Println(total)
			hasNextPage = response.Search.PageInfo.HasNextPage
			nextPageCursor = response.Search.PageInfo.EndCursor
		}
	}
	producer.Flush(15 * 1000)
	producer.Close()
}

//go:generate go run github.com/Khan/genqlient
