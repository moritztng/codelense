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
	minCreated := "1970-01-01"
	createdAfterMin := true
	var lastKeys map[int]struct{}
	total := 0
	for createdAfterMin {
		keys := map[int]struct{}{}
		createdAfterMin = false
		hasNextPage := true
		nextPageCursor := ""
		query := fmt.Sprintf("location:Germany type:org sort:joined-asc created:>=%s", minCreated)
		for hasNextPage {
			response, _ := getOrganizations(context.Background(), graphqlClient, nextPageCursor, query)
			edges := response.Search.Edges
			for _, edge := range edges {
				organization := edge.Node.(*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization)
				_, exists := lastKeys[organization.DatabaseId]
				if exists {
					continue
				}
				organizationJson, _ := json.Marshal(model.Organization{GithubID: uint(organization.DatabaseId), Login: organization.Login, Name: organization.Name, Location: organization.Location, GithubCreatedAt: organization.CreatedAt, Email: organization.Email, AvatarUrl: organization.AvatarUrl, Description: organization.Description, TwitterUsername: organization.TwitterUsername, GithubUpdatedAt: organization.UpdatedAt, WebsiteUrl: organization.WebsiteUrl, Url: organization.Url})
				topic := "github_load_organizations"
				kafkaProducer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          organizationJson,
				}, nil)
				created := organization.CreatedAt.Format("2006-01-02")
				if created != minCreated {
					minCreated = created
					createdAfterMin = true
				}
				keys[organization.DatabaseId] = struct{}{}
				total += 1
			}
			logger.Infow("loaded organizations", "created_after", minCreated, "total", total)
			hasNextPage = response.Search.PageInfo.HasNextPage
			nextPageCursor = response.Search.PageInfo.EndCursor
		}
		lastKeys = keys
	}
	kafkaProducer.Flush(15 * 1000)
	logger.Info("stop")
	kafkaProducer.Close()
}

//go:generate go run github.com/Khan/genqlient
