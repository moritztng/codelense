package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/messaging"
)

type Repo struct {
	Name  string
	Stars uint
}
type githubStoreResult struct {
	Name  string
	Stars uint
}

type apiGithubRequest struct {
	Name string
}

func main() {
	conn, _ := pgx.Connect(context.Background(), "postgresql://postgres:example@localhost:5001/postgres")
	defer conn.Close(context.Background())
	conf := messaging.ReadConfig("kafka.properties")
	consumer, _ := kafka.NewConsumer(&conf)
	producer, _ := kafka.NewProducer(&conf)
	defer producer.Close()
	defer consumer.Close()
	consumer.SubscribeTopics([]string{"github", "api_github_requests"}, nil)
	for {
		message, err := consumer.ReadMessage(time.Second)
		if err == nil {
			switch *message.TopicPartition.Topic {
			case "github":
				var repo Repo
				json.Unmarshal(message.Value, &repo)
				tx, _ := conn.Begin(context.Background())
				defer tx.Rollback(context.Background())
				_, err = tx.Exec(context.Background(), "insert into github(name, stars) values ($1,$2)", repo.Name, repo.Stars)
				if err != nil {
					fmt.Println(err)
					return
				}
				err = tx.Commit(context.Background())
				if err != nil {
					fmt.Println(err)
					return
				}
			case "api_github_requests":
				fmt.Println("github_request")
				var request apiGithubRequest
				json.Unmarshal(message.Value, &request)
				result, _ := json.Marshal(githubStoreResult{request.Name, 2})
				topic := "github_store_results"
				fmt.Println("produce")
				producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          result,
				}, nil)
				fmt.Println("end produce")

			}
		} else {
			fmt.Println(err)
		}
	}
}
