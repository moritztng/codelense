package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_TIMEZONE"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(&model.Event{}, &model.OrganizationEvent{})
	kafkaConfig := kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "github_store",
		"auto.offset.reset": "earliest",
	}
	kafkaConsumer, _ := kafka.NewConsumer(&kafkaConfig)
	defer kafkaConsumer.Close()
	kafkaConsumer.SubscribeTopics([]string{"github_load_organizations", "github_load_events"}, nil)
	for {
		message, err := kafkaConsumer.ReadMessage(time.Second)
		if err == nil {
			switch *message.TopicPartition.Topic {
			case "github_load_organizations":
				var organizationEvent model.OrganizationEvent
				json.Unmarshal(message.Value, &organizationEvent)
				database.Create(&organizationEvent)
			case "github_load_events":
				var event model.Event
				json.Unmarshal(message.Value, &event)
				database.Create(&event)
			}
		} else {
			fmt.Println(err)
		}
	}
}
