package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/messaging"
)

func main() {
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	databaseConn, _ := pgx.Connect(context.Background(), databaseUrl)
	defer databaseConn.Close(context.Background())
	conf := messaging.ReadConfig("kafka.properties")
	consumer, _ := kafka.NewConsumer(&conf)
	producer, _ := kafka.NewProducer(&conf)
	defer producer.Close()
	defer consumer.Close()
	consumer.SubscribeTopics([]string{"github_load_organizations", "api_github_requests"}, nil)
	for {
		message, err := consumer.ReadMessage(time.Second)
		if err == nil {
			switch *message.TopicPartition.Topic {
			case "github_load_organizations":
				var organization messaging.Organization
				json.Unmarshal(message.Value, &organization)
				tx, _ := databaseConn.Begin(context.Background())
				defer tx.Rollback(context.Background())
				_, err = tx.Exec(context.Background(), "insert into organizations(key, login, name, created) values ($1, $2, $3, $4)", organization.Key, organization.Login, organization.Name, organization.CreatedAt)
				fmt.Println(err)
				err = tx.Commit(context.Background())
				fmt.Println(err)
				/*case "api_github_requests":
				var request messaging.ApiGithubRequest
				json.Unmarshal(message.Value, &request)
				rows, _ := databaseConn.Query(context.Background(), "select key, login, name, created from organizations where stars<=$1 order by stars desc limit $2", request.MaxStars, request.First)
				defer rows.Close()
				repositories, _ := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[messaging.Repository])
				result, _ := json.Marshal(messaging.GithubStoreResult{Key: request.Key, Repositories: repositories})
				produceTopic := "github_store_results"
				producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &produceTopic, Partition: kafka.PartitionAny},
					Value:          result,
				}, nil)*/
			}
		} else {
			fmt.Println(err)
		}
	}
}
