package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/messaging"
	"github.com/moritztng/codelense/backend/services/api/graph"
)

func main() {
	conf := messaging.ReadConfig("kafka.properties")
	producer, _ := kafka.NewProducer(&conf)
	consumer, _ := kafka.NewConsumer(&conf)
	consumer.SubscribeTopics([]string{"github_store_results"}, nil)
	defer producer.Close()
	defer consumer.Close()

	port := os.Getenv("PORT")
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Consumer: consumer, Producer: producer}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
