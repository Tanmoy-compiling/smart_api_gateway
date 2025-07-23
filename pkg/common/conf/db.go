package conf

import (
    "context"
    "fmt"
    "os"
    "log"

    "github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDB() error {
    connStr := os.Getenv("DATABASE_URL")
    if connStr == "" {
        return fmt.Errorf("DATABASE_URL is not set in environment")
    }

    var err error
    DB, err = pgx.Connect(context.Background(), connStr)
    if err != nil {
        return fmt.Errorf("failed to connect to DB: %w", err)
    }

    log.Println("✅ Connected to Neon/Postgres DB")
    return nil
}

func CloseDB() {
    if DB != nil {
        DB.Close(context.Background())
        log.Println("DB connection closed")
    }
}
