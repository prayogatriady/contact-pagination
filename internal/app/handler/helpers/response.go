package helpers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func StatusOK(message string) map[string]any {
	return map[string]any{
		"code":    http.StatusOK,
		"message": message,
	}
}

func StatusOKWithData(message string, data interface{}) map[string]any {
	return map[string]any{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	}
}

func InternalServerError() map[string]any {
	return map[string]any{
		"status": http.StatusInternalServerError,
		"message": "Something went wrong on the server side",
	}
}

func StatusBadRequest(err error) map[string]any {
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = e.Tag()
	}

	return map[string]any{
		"status": http.StatusBadRequest,
		"message": "Something went wrong on the request",
		"errors": errors,
	}
}

func StatusBadRequest_Query(query string) map[string]any {
	return map[string]any{
		"status": http.StatusBadRequest,
		"message": fmt.Sprintf("Something went wrong on the request query: %s", query),
	}
}

func StatusBadRequest_Param(param string) map[string]any {
	return map[string]any{
		"status": http.StatusBadRequest,
		"message": fmt.Sprintf("Something went wrong on the request param: %s", param),
	}
}