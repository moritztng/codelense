package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-co-op/gocron"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

func main() {
	baseLogger, _ := zap.NewProduction()
	logger := baseLogger.Sugar()
	defer logger.Sync()
	logger.Info("start")
	producer, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))})
	defer producer.Close()
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Cron(os.Getenv("CRON")).Do(func() {
		logger.Infow("schedule", "name", os.Getenv("NAME"))
		produceTopic := fmt.Sprintf("schedule_%s", os.Getenv("NAME"))
		fmt.Print(os.Getenv("MESSAGE"))
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &produceTopic, Partition: kafka.PartitionAny},
			Value:          []byte(os.Getenv("MESSAGE")),
		}, nil)
	})
	scheduler.StartBlocking()
	logger.Info("stop")
}
