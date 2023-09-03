package service

import (
	"math"
	"os"
	"testing"

	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var contactRepositoryMock *repository.ContactRepositoryMock
var contactServ ContactService

func TestMain(m *testing.M) {

	// Create a mock repository
	contactRepositoryMock = &repository.ContactRepositoryMock{}
	contactServ = NewContactService(contactRepositoryMock)

	result := m.Run()

	os.Exit(result)
	
}

func TestContactService_GetContactList(t *testing.T) {

	// Set up parameter data
	pagination := &model.Pagination{
		Limit: 2,
		Page:  1,
		Sort:  "Id asc",
	}

	// Define the expected data
	request := &model.PaginationRequest{
		Limit: pagination.Limit,
		Page:  pagination.Page,
		Sort:  pagination.Sort,
	}
	expectedContacts := []*model.Contact{
		{ID: 1, FirstName: "Prayoga", LastName: "Triady"},
		{ID: 2, FirstName: "James", LastName: "Jumes"},
		{ID: 3, FirstName: "Fullan", LastName: "Fullin"},
		{ID: 4, FirstName: "Aisyah"},
		{ID: 5, FirstName: "Ali"},
	}
	expectedTotalRows := int64(len(expectedContacts))
	expectedTotalPages := int(math.Ceil(float64(expectedTotalRows) / float64(pagination.GetLimit())))

	// Configure the mock repository behavior
	contactRepositoryMock.On("GetContactList", pagination).Return(expectedContacts[:pagination.Limit], expectedTotalRows, nil)

	// Call the method being tested
	response, err := contactServ.GetContactList(request)
	// fmt.Printf("%+v\n", response)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedTotalRows, response.TotalRows)
	assert.Equal(t, expectedTotalPages, response.TotalPages)
	assert.Equal(t, pagination.GetLimit(), len(response.Rows))

	// Verify that expected methods were called on the mock repository
	contactRepositoryMock.AssertCalled(t, "GetContactList", pagination)
}

func TestContactService_GetContact(t *testing.T) {

	// Set up parameter data
	contactId := 1

	// Define the expected data
	expectedContacts := &model.Contact{
		ID: 1, 
		FirstName: "Prayoga", 
		LastName: "Triady",
		Email: "yoga@gmail.com",
	}

	// Configure the mock repository behavior
	contactRepositoryMock.On("GetContact", contactId).Return(expectedContacts, nil)

	// Call the method being tested
	response, err := contactServ.GetContact(contactId)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 1, response.ID)
	assert.Equal(t, "Prayoga", response.FirstName)
	assert.Empty(t, response.Phone)

	// Verify that expected methods were called on the mock repository
	contactRepositoryMock.AssertCalled(t, "GetContact", contactId)
}

func TestContactService_CreateContact(t *testing.T) {

	// Set up parameter data
	request := &model.CreateContactRequest{
		FirstName: "Prayoga", 
		LastName: "Triady",
		Email: "yoga@gmail.com",
		Phone: "081211112222",
	}

	// Define the expected data
	expectedContact := &model.Contact{
		FirstName: "Prayoga",
		LastName:  "Triady",
		Email:     "yoga@gmail.com",
		Phone:     "081211112222",
	}

	// Configure the mock repository behavior
    contactRepositoryMock.On("CreateContact", mock.AnythingOfType("*model.Contact")).Return(nil).Run(func(args mock.Arguments) {
        // Here, you can access the created contact from the argument and assert it
        createdContact := args.Get(0).(*model.Contact)
        assert.Equal(t, expectedContact.FirstName, createdContact.FirstName)
        assert.Equal(t, expectedContact.LastName, createdContact.LastName)
        assert.Equal(t, expectedContact.Email, createdContact.Email)
        assert.Equal(t, expectedContact.Phone, createdContact.Phone)
    })

	// Call the method being tested
	err := contactServ.CreateContact(request)
	// Assertions
	assert.NoError(t, err)

	// Verify that expected methods were called on the mock repository
	contactRepositoryMock.AssertCalled(t, "CreateContact", mock.AnythingOfType("*model.Contact"))
}

func TestContactService_UpdateContact(t *testing.T) {

	// Set up parameter data
	contactId := 1

	request := &model.UpdateContactRequest{
		FirstName: "Prayoga", 
		LastName: "Triady",
		Email: "yoga@gmail.com",
		Phone: "081211112222",
	}

	// Define the expected data
	expectedContact := &model.Contact{
		FirstName: "Prayoga",
		LastName:  "Triady",
		Email:     "yoga@gmail.com",
		Phone:     "081211112222",
	}

	// Configure the mock repository behavior
    contactRepositoryMock.On("UpdateContact", mock.AnythingOfType("*model.Contact")).Return(nil).Run(func(args mock.Arguments) {
        // Here, you can access the created contact from the argument and assert it
        updatedContact := args.Get(0).(*model.Contact)
        assert.Equal(t, expectedContact.FirstName, updatedContact.FirstName)
        assert.Equal(t, expectedContact.LastName, updatedContact.LastName)
        assert.Equal(t, expectedContact.Email, updatedContact.Email)
        assert.Equal(t, expectedContact.Phone, updatedContact.Phone)
    })

	// Call the method being tested
	err := contactServ.UpdateContact(contactId, request)
	// Assertions
	assert.NoError(t, err)

	// Verify that expected methods were called on the mock repository
	contactRepositoryMock.AssertCalled(t, "UpdateContact", mock.AnythingOfType("*model.Contact"))
}