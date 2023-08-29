package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/service"
)

func TestContactHandler_Paginate(t *testing.T) {
	// Create a mock ContactService
	mockService := &service.ContactServiceMock{}
	handler := NewContactHandler(mockService)

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
		// ...
	}

	// Configure the mock ContactService behavior
	mockService.On("Paginate", expectedRequest).Return(expectedResponse, nil)

	// Call the method being tested
	handler.Paginate(c)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify that expected methods were called on the mock ContactService
	mockService.AssertCalled(t, "Paginate", expectedRequest)
}
