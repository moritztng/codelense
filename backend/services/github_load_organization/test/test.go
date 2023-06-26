package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/model"
)

func main() {
	producer, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))})
	defer producer.Close()
	produceTopic := "github_organization_events"
	value, _ := json.Marshal(model.OrganizationEvent{Login: "NVIDIA"})
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &produceTopic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)
	producer.Flush(15 * 1000)
}
