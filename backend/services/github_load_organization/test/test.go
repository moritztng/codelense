package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	producer, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))})
	defer producer.Close()
	produceTopic := "organization_events"
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &produceTopic, Partition: kafka.PartitionAny},
		Value:          []byte("NVIDIA"),
	}, nil)
	producer.Flush(15 * 1000)
}
