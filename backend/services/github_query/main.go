package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	baseLogger, _ := zap.NewProduction()
	logger := baseLogger.Sugar()
	defer logger.Sync()
	logger.Info("start")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_TIMEZONE"))
	database, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	kafkaConfig := kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
		"group.id":          "github_filter_events_org",
		"auto.offset.reset": "earliest",
	}
	kafkaConsumer, _ := kafka.NewConsumer(&kafkaConfig)
	defer kafkaConsumer.Close()
	kafkaConsumer.SubscribeTopics([]string{"schedule_github_query"}, nil)
	for {
		message, err := kafkaConsumer.ReadMessage(time.Second)
		if err != nil {
			continue
		}
		var query struct {
			Name  string `json:"queryName"`
			Query string `json:"query"`
		}
		json.Unmarshal(message.Value, &query)
		database.Exec(fmt.Sprintf("CREATE MATERIALIZED VIEW IF NOT EXISTS %s AS %s", query.Name, query.Query))
		database.Exec(fmt.Sprintf("REFRESH MATERIALIZED VIEW %s", query.Name))
		logger.Infow("query", "name", query.Name, "query", query.Query)
	}
	logger.Info("stop")
}
