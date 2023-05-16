package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:example@localhost:5001/postgres")
	return conn, err
}
