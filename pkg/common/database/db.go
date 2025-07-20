package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	config := &Config{
		Host:     os.Getenv("NEON_HOST"),
		Port:     os.Getenv("NEON_PORT"),
		User:     os.Getenv("NEON_USER"),
		Password: os.Getenv("NEON_PASSWORD"),
		DBName:   os.Getenv("NEON_DB"),
		SSLMode:  "require",
	}

	if config.Port == "" {
		config.Port = "5432"
	}

	var err error
	DB, err = NewConnection(config)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("✅ Connected to Neon DB")
	return nil
}

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}
