package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/service"
	"github.com/stretchr/testify/mock"
)

var mockService *service.ContactServiceMock
var handler ContactHandler

func TestMain(m *testing.M) {

	// Create a mock ContactService
	mockService = &service.ContactServiceMock{}
	handler = NewContactHandler(mockService)

	result := m.Run()

	os.Exit(result)
}

func TestContactHandler_GetContactList(t *testing.T) {

	// Create a mock HTTP request and response recorder
	req, _ := http.NewRequest(http.MethodGet, "/api/contacts?limit=10&page=1&sort=Id asc", nil)
	w := httptest.NewRecorder()

	// Create a Gin context manually using the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set up expected input and output
	expectedRequest := &model.PaginationRequest{
		Limit: 10,
		Page:  1,
		Sort:  "Id asc",
	}
	expectedResponse := &model.PaginationResponse{
		Limit: 10,
		Page:  1,
		Sort:  "Id asc",
	}

	// Configure the mock ContactService behavior
	mockService.On("GetContactList", expectedRequest).Return(expectedResponse, nil)

	// Call the method being tested
	handler.GetContactList(c)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify that expected methods were called on the mock ContactService
	mockService.AssertCalled(t, "GetContactList", expectedRequest)
}

func TestContactHandler_GetContactList_BadLimit(t *testing.T) {

	// SKIP TEST
	t.Skip()

	// Create a mock HTTP request with an incorrect URL
	req, _ := http.NewRequest(http.MethodGet, "/api/contacts?limit=non-numeric&page=1&sort=Id asc", nil)
	w := httptest.NewRecorder()

	// Create a Gin context manually using the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Configure the mock ContactService behavior
	mockService.On("GetContactList",  mock.Anything).Return( mock.Anything, nil)

	// Call the method being tested
	handler.GetContactList(c)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Verify that the mock ContactService was not called
	mockService.AssertNotCalled(t, "GetContactList", mock.Anything)
}

func TestContactHandler_GetContactList_BadSort(t *testing.T) {

	// SKIP TEST
	t.Skip()

	// Create a mock HTTP request with an incorrect URL
	req, _ := http.NewRequest(http.MethodGet, "/api/contacts?limit=2&page=1&sort=Id asc,", nil)
	w := httptest.NewRecorder()

	// Create a Gin context manually using the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Configure the mock ContactService behavior
	mockService.On("GetContactList",  mock.Anything).Return( mock.Anything, nil)

	// Call the method being tested
	handler.GetContactList(c)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Verify that the mock ContactService was not called
	mockService.AssertNotCalled(t, "GetContactList", mock.Anything)
}

func TestContactHandler_GetContact(t *testing.T) {

	// SKIP TEST
	t.Skip()

	// Create a mock HTTP request and response recorder
	req, _ := http.NewRequest(http.MethodGet, "/api/contact/1", nil)
	w := httptest.NewRecorder()

	// Create a Gin context manually using the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set up expected input and output
	contactId := 1

	expectedContacts := &model.ContactResponse{
		ID: 1, 
		FirstName: "Prayoga", 
		LastName: "Triady",
		Email: "yoga@gmail.com",
		Phone: "081122223333",
	}

	// Configure the mock ContactService behavior
	mockService.On("GetContact", contactId).Return(expectedContacts, nil)

	// Call the method being tested
	handler.GetContact(c)

	fmt.Printf("Actual contactId: %d\n", contactId)

	// Verify that expected methods were called on the mock ContactService
	mockService.AssertCalled(t, "GetContact", contactId)
}

func TestContactHandler_CreateContact(t *testing.T) {

	// Create a JSON request body
	requestBody := `{
		"first_name": "John",
		"last_name": "Doe",
		"email": "yoga@gmail.com",
		"phone": "081122223333"
	}`

	// Create a mock HTTP request and response recorder
	req, _ := http.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Create a Gin context manually using the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set up expected input and output
	expectedRequest := &model.CreateContactRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "yoga@gmail.com",
		Phone:     "081122223333",
	}
	// expectedError := nil // Replace with the expected error value if needed

	// Configure the mock ContactService behavior
	mockService.On("CreateContact", expectedRequest).Return(nil)

	// Call the method being tested
	handler.CreateContact(c)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify that expected methods were called on the mock ContactService
	mockService.AssertCalled(t, "CreateContact", expectedRequest)
}

func TestContactHandler_UpdateContact(t *testing.T) {

	// SKIP TEST
	t.Skip()

	// Create a JSON request body
	requestBody := `{
		"first_name": "John",
		"last_name": "Doe",
		"email": "yoga@gmail.com",
		"phone": "081122223333"
	}`

	// Create a mock HTTP request and response recorder
	req, _ := http.NewRequest(http.MethodPut, "/api/contact/1", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Create a Gin context manually using the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set up expected input and output
	contactId := 1

	expectedRequest := &model.UpdateContactRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "yoga@gmail.com",
		Phone:     "081122223333",
	}
	// expectedError := nil // Replace with the expected error value if needed

	// Configure the mock ContactService behavior
	mockService.On("UpdateContact", contactId, expectedRequest).Return(nil)

	// Call the method being tested
	handler.UpdateContact(c)

	// Assertions
	// assert.Equal(t, http.StatusOK, w.Code)

	// Verify that expected methods were called on the mock ContactService
	mockService.AssertCalled(t, "UpdateContact", contactId, expectedRequest)
}

func TestContactHandler_DeleteContact(t *testing.T) {

	// SKIP TEST
	t.Skip()

	// Create a mock HTTP request and response recorder
	req, _ := http.NewRequest(http.MethodDelete, "/api/contact/1", nil)
	w := httptest.NewRecorder()

	// Create a Gin context manually using the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set up expected input and output
	contactId := 1

	// Configure the mock ContactService behavior
	mockService.On("DeleteContact", contactId).Return(nil)

	// Call the method being tested
	handler.DeleteContact(c)

	// Assertions
	// assert.Equal(t, http.StatusOK, w.Code)

	// Verify that expected methods were called on the mock ContactService
	mockService.AssertCalled(t, "DeleteContact", contactId)
}