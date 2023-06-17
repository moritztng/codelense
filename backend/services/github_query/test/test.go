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
	produceTopic := "schedule_github_query"
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &produceTopic, Partition: kafka.PartitionAny},
		Value:          []byte("{\"queryName\": \"repositories_time_stars\", \"query\": \"SELECT to_timestamp(round(extract('epoch' from github_created_at) / 3600) * 3600) as time, repository_name, count(*) as star_count FROM events WHERE type='WatchEvent' GROUP BY time, repository_name HAVING count(*)>=10 ORDER BY time\"}"),
	}, nil)
	producer.Flush(15 * 1000)
}
