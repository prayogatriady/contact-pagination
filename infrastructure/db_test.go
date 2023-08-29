package infrastructure

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_NewDatabase(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.NoError(t, err)

	_, err = NewDatabase(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	assert.NoError(t, err)
}
