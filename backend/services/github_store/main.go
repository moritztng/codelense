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
	"github.com/moritztng/codelense/backend/util"
)

func main() {
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	databaseConn, _ := pgx.Connect(context.Background(), databaseUrl)
	defer databaseConn.Close(context.Background())
	conf := util.ReadConfig("kafka.properties")
	consumer, _ := kafka.NewConsumer(&conf)
	producer, _ := kafka.NewProducer(&conf)
	defer producer.Close()
	defer consumer.Close()
	consumer.SubscribeTopics([]string{"github_load_organizations", "github_load_events"}, nil)
	for {
		message, err := consumer.ReadMessage(time.Second)
		if err == nil {
			switch *message.TopicPartition.Topic {
			case "github_load_organizations":
				var organization util.Organization
				json.Unmarshal(message.Value, &organization)
				tx, _ := databaseConn.Begin(context.Background())
				defer tx.Rollback(context.Background())
				_, err = tx.Exec(context.Background(), "insert into organizations(key, login, name, email, description, members_count, repositories_count, twitter_username, website_url, url, avatar_url, created, updated) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", organization.Key, organization.Login, organization.Name, organization.Email, organization.Description, organization.MembersCount, organization.RepositoriesCount, organization.TwitterUsername, organization.WebsiteUrl, organization.Url, organization.AvatarUrl, organization.CreatedAt, organization.UpdatedAt)
				fmt.Println(err)
				err = tx.Commit(context.Background())
				fmt.Println(err)
			case "github_load_events":
				var event util.Event
				json.Unmarshal(message.Value, &event)
				tx, _ := databaseConn.Begin(context.Background())
				defer tx.Rollback(context.Background())
				_, err = tx.Exec(context.Background(), "insert into events(key, type, actor_id, org_id, repository_id, payload, public, created_at) values ($1, $2, $3, $4, $5, $6, $7, $8)", event.Key, event.Type, event.ActorId, event.OrgId, event.RepositoryId, event.Payload, event.Public, event.CreatedAt)
				fmt.Println(err)
				err = tx.Commit(context.Background())
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}
}
