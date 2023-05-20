package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/moritztng/codelense/backend/messaging"
	"github.com/moritztng/codelense/backend/services/api/graph"
)

const defaultPort = "8080"

type Repo struct {
	Name  string
	Stars uint
}

func main() {
	conf := messaging.ReadConfig("kafka.properties")
	producer, _ := kafka.NewProducer(&conf)
	consumer, _ := kafka.NewConsumer(&conf)
	consumer.SubscribeTopics([]string{"github_store_results"}, nil)
	defer producer.Close()
	defer consumer.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Consumer: consumer, Producer: producer}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
