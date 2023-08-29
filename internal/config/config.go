package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration values
type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	Port       string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
	}

	return config, nil
}
