package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-co-op/gocron"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	producer, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))})
	defer producer.Close()
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Cron(os.Getenv("CRON")).Do(func() {
		produceTopic := fmt.Sprintf("scheduler_%s", os.Getenv("NAME"))
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &produceTopic, Partition: kafka.PartitionAny},
		}, nil)
	})
	scheduler.StartBlocking()
}
