package service

import (
	"fmt"
	"testing"

	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestContactService_Paginate(t *testing.T) {
	// Create a mock repository
	contactRepositoryMock := &repository.ContactRepositoryMock{}
	contactService := NewContactService(contactRepositoryMock)

	// Set up expected data
	request := &model.PaginationRequest{
		Limit: 2,
		Page:  1,
		Sort:  "Id asc",
	}
	expectedContacts := []*model.Contact{
		{FirstName: "Prayoga", LastName: "Triady"},
		{FirstName: "James", LastName: "Jumes"},
	}
	expectedContactsResponse := []*model.ContactResponse{
		{FirstName: "Prayoga", LastName: "Triady"},
		{FirstName: "James", LastName: "Jumes"},
	}
	expectedTotalRows := int64(len(expectedContacts))
	expectedTotalPages := 1

	// Configure the mock repository behavior
	contactRepositoryMock.On("Paginate", mock.Anything, mock.Anything, mock.Anything).Return(expectedContacts, nil)
	contactRepositoryMock.On("Count", expectedContacts).Return(expectedTotalRows, nil)

	// Call the method being tested
	response, err := contactService.Paginate(request)
	fmt.Println(response)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedTotalRows, response.TotalRows)
	assert.Equal(t, expectedTotalPages, response.TotalPages)
	assert.Equal(t, expectedContactsResponse, response.Rows)

	// Verify that expected methods were called on the mock repository
	contactRepositoryMock.AssertCalled(t, "Paginate", mock.Anything, mock.Anything, mock.Anything)
	contactRepositoryMock.AssertCalled(t, "Count", expectedContacts)
}
