package graph

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Database           *pgxpool.Pool
	Logger             *zap.SugaredLogger
	EventsQuery        string
	OrganizationsQuery string
}
