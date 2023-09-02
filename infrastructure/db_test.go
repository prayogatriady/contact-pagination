package infrastructure

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/prayogatriady/golang-rest-pagination/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_NewDatabase(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.NoError(t, err)

	var cfg config.AppConfig
	cfg.Database.DBUser = os.Getenv("DB_USER")
	cfg.Database.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.Database.DBHost = os.Getenv("DB_HOST")
	cfg.Database.DBPort = os.Getenv("DB_PORT")
	cfg.Database.DBName = os.Getenv("DB_NAME")

	fmt.Println(cfg.Database.DBName)

	_, err = NewDatabase(&cfg)
	assert.NoError(t, err)
}
