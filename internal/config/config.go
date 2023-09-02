package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// Config holds application configuration values
type AppConfig struct {

	App struct {
		BaseURL  string
		Port     string
		ENV      string
	}

	Database struct {
		DBUser     string
		DBPassword string
		DBHost     string
		DBPort     string
		DBName     string
	}
}

var m sync.Mutex
var appConfig *AppConfig

func NewConfig() *AppConfig {
	m.Lock()
	defer m.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {

	var config AppConfig

	config.App.BaseURL = GetEnv("APP_PORT", "http://localhost:8000")
	config.App.Port = GetEnv("APP_BASE_URL", "8000")
	config.App.ENV = GetEnv("APP_ENV", "development")

	config.Database.DBUser = GetEnv("DB_USER", "")
	config.Database.DBPassword = GetEnv("DB_PASSWORD", "")
	config.Database.DBHost = GetEnv("DB_HOST", "")
	config.Database.DBPort = GetEnv("DB_PORT", "")
	config.Database.DBName = GetEnv("DB_NAME", "")

	return &config
}

func GetEnv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}