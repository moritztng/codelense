package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/services/api/graph"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

func main() {
	baseLogger, _ := zap.NewProduction()
	logger := baseLogger.Sugar()
	defer logger.Sync()
	logger.Info("start")
	pgxPool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		logger.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pgxPool.Close()

	eventsQuery, _ := os.ReadFile("query_events.sql")
	organizationsQuery, _ := os.ReadFile("query_organizations.sql")

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Database: pgxPool, Logger: logger, EventsQuery: string(eventsQuery), OrganizationsQuery: string(organizationsQuery)}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	logger.Fatal(http.ListenAndServe(":"+os.Getenv("API_PORT"), router))
	logger.Info("stop")
}
