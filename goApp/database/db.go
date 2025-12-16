package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return fmt.Errorf("DB_URL not set")
	}

	var err error
	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return err
	}

	// ✅ test connection
	return DB.Ping(context.Background())
}
