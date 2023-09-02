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

// Set up the testing database environment
func setUpDB() (*gorm.DB, error) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		return nil, err
	}

	var cfg config.AppConfig
	cfg.Database.DBUser = os.Getenv("DB_USER_TEST")
	cfg.Database.DBPassword = os.Getenv("DB_PASSWORD_TEST")
	cfg.Database.DBHost = os.Getenv("DB_HOST_TEST")
	cfg.Database.DBPort = os.Getenv("DB_PORT_TEST")
	cfg.Database.DBName = os.Getenv("DB_NAME_TEST")

	db, err := infrastructure.NewDatabase(&cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Initialize list of contact data for all unit test
func getContactsData() []*model.Contact {
	return []*model.Contact{
		{ID: 1, FirstName: "Prayoga", LastName: "Triady"},
		{ID: 2, FirstName: "James", LastName: "Jumes"},
		{ID: 3, FirstName: "Fullan", LastName: "Fullin"},
		{ID: 4, FirstName: "Aisyah"},
		{ID: 5, FirstName: "Ali"},
	}
}

// Insert testing data to the testing database
func createTestData(db *gorm.DB) error {
	contacts := getContactsData()
	return db.Create(&contacts).Error
}

// Delete testing data from the testing database
func cleanupTestData(db *gorm.DB) error {
	return db.Exec("DELETE FROM contacts").Error
}

// Unit test function for testing GetContactList
func TestContactRepository_GetContactList(t *testing.T) {
	db, err := setUpDB()
	if err != nil {
		log.Fatal("Error initialize database:", err)
	}

	assert.NoError(t, createTestData(db))

	repo := NewUserRepository(db)
	pagination := model.Pagination{
		Limit: 2,
		Page: 2,
		Sort: "Id asc",
	}
	contacts, totalRows, err := repo.GetContactList(&pagination)
	assert.NoError(t, err)

	assert.Equal(t, 3, contacts[0].ID)
	assert.Equal(t, 4, contacts[1].ID)
	assert.Equal(t, int64(5), totalRows)

	assert.NoError(t, cleanupTestData(db))
}

// Unit test function for testing GetContact
func TestContactRepository_GetContact(t *testing.T) {
	db, err := setUpDB()
	if err != nil {
		log.Fatal("Error initialize database:", err)
	}

	assert.NoError(t, createTestData(db))

	repo := NewUserRepository(db)
	contact, err := repo.GetContact(1)
	assert.NoError(t, err)

	assert.Equal(t, "Prayoga", contact.FirstName)
	assert.Equal(t, "Triady", contact.LastName)

	assert.NoError(t, cleanupTestData(db))
}

// Unit test function for testing CreateContact
func TestContactRepository_CreateContact(t *testing.T) {
	db, err := setUpDB()
	if err != nil {
		log.Fatal("Error initialize database:", err)
	}

	contacts := getContactsData()

	repo := NewUserRepository(db)
	assert.NoError(t, repo.CreateContact(contacts[0]))

	contactDb, _ := repo.GetContact(1)
	assert.Equal(t, "Prayoga", contactDb.FirstName)
	assert.Equal(t, "Triady", contactDb.LastName)

	assert.NoError(t, cleanupTestData(db))
}

// Unit test function for testing UpdateContact
func TestContactRepository_UpdateContact(t *testing.T) {
	db, err := setUpDB()
	if err != nil {
		log.Fatal("Error initialize database:", err)
	}

	assert.NoError(t, createTestData(db))

	repo := NewUserRepository(db)
	contact := model.Contact{
		ID: 1, 
		// FirstName: "Prayoga", 
		// LastName: "Triady",
		Email: "yoga@gmail.com",
		Phone: "081211112222",
	}
	assert.NoError(t, repo.UpdateContact(&contact))

	contactDb, err := repo.GetContact(1)
	assert.Equal(t, "Prayoga", contactDb.FirstName)
	assert.Equal(t, "Triady", contactDb.LastName)
	assert.Equal(t, "yoga@gmail.com", contactDb.Email)
	assert.Equal(t, "081211112222", contactDb.Phone)

	assert.NoError(t, cleanupTestData(db))
}

// Unit test function for testing DeleteContact
func TestContactRepository_DeleteContact(t *testing.T) {
	db, err := setUpDB()
	if err != nil {
		log.Fatal("Error initialize database:", err)
	}

	assert.NoError(t, createTestData(db))

	repo := NewUserRepository(db)

	assert.NoError(t, repo.DeleteContact(1))

	contactDb, err := repo.GetContact(1)
	assert.NotNil(t, contactDb.DeletedAt)

	assert.NoError(t, cleanupTestData(db))
}
