package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func LoadEnv(path string) error {
	if err := godotenv.Load(path); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}
	return nil
}

func NewConfig() (*Config, error) {
	key := os.Getenv("KEY")
	if key == "" {
		return nil, fmt.Errorf("empty KEY in environment variables")
	}

	return &Config{
		Key: key,
	}, nil
}
