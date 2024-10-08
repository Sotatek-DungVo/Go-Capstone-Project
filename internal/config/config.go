package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
	LogLevel      string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		LogLevel:      os.Getenv("LOG_LEVEL"),
	}

	// Set default values if not provided
	if config.DatabaseURL == "" {
		config.DatabaseURL = "sqlite://ecombase.db"
	}
	if config.ServerAddress == "" {
		config.ServerAddress = ":8080"
	}
	if config.LogLevel == "" {
		config.LogLevel = "info"
	}

	return config, nil
}
