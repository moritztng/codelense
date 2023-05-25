package main

import (
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-co-op/gocron"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/messaging"
)

func main() {
	conf := messaging.ReadConfig("kafka.properties")
	producer, _ := kafka.NewProducer(&conf)
	defer producer.Close()
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Cron(os.Getenv("CRON")).Do(func() {
		produceTopic := "github_load_scheduler"
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &produceTopic, Partition: kafka.PartitionAny},
		}, nil)
	})
	scheduler.StartBlocking()
}
