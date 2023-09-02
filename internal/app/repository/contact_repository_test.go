package repository

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/prayogatriady/golang-rest-pagination/infrastructure"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/prayogatriady/golang-rest-pagination/internal/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setUpDB() (*gorm.DB, error) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		return nil, err
	}

	var cfg config.AppConfig
	cfg.Database.DBUser = os.Getenv("DB_USER")
	cfg.Database.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.Database.DBHost = os.Getenv("DB_HOST")
	cfg.Database.DBPort = os.Getenv("DB_PORT")
	cfg.Database.DBName = os.Getenv("DB_NAME")

	db, err := infrastructure.NewDatabase(&cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTestData(db *gorm.DB) error {
	contacts := []*model.Contact{
		{FirstName: "Prayoga", LastName: "Triady"},
		{FirstName: "James", LastName: "Jumes"},
		{FirstName: "Fullan", LastName: "Fullin"},
		{FirstName: "Aisyah"},
		{FirstName: "Ali"},
	}

	return db.Create(&contacts).Error
}

func cleanupTestData(db *gorm.DB) error {
	return db.Exec("DELETE FROM contacts").Error
}

func TestContactRepository_Paginate(t *testing.T) {
	db, err := setUpDB()
	if err != nil {
		log.Fatal("Error initialize database:", err)
	}

	assert.NoError(t, createTestData(db))

	repo := NewUserRepository(db)
	contacts, err := repo.Paginate(1, 2, "Id asc")
	assert.NoError(t, err)

	assert.Equal(t, "James", contacts[0].FirstName)
	assert.Equal(t, "Fullan", contacts[1].FirstName)

	assert.NoError(t, cleanupTestData(db))
}

func TestContactRepository_Count(t *testing.T) {
	db, err := setUpDB()
	if err != nil {
		log.Fatal("Error initialize database:", err)
	}

	assert.NoError(t, createTestData(db))

	repo := NewUserRepository(db)
	totalRows, err := repo.Count(model.Contact{})
	assert.NoError(t, err)

	assert.Equal(t, int64(5), totalRows)

	assert.NoError(t, cleanupTestData(db))
}
