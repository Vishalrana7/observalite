package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	PostgresURL string
}

func LoadConfigFromEnv() Config {
	return Config{
		PostgresURL: os.Getenv("POSTGRES_URL"),
	}
}

func Connect(cfg Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), cfg.PostgresURL)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return conn, nil
}
