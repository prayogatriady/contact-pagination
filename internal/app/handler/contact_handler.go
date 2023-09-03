package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/service"
)

type ContactHandler interface {
	Paginate(c *gin.Context)
}

type contactHandler struct {
	contactService service.ContactService
}

func NewContactHandler(contactService service.ContactService) ContactHandler {
	return &contactHandler{
		contactService: contactService,
	}
}

func (handler *contactHandler) Paginate(c *gin.Context) {
	var (
		request  *model.PaginationRequest
		limit    int
		page     int
		err      error
		response *model.PaginationResponse
	)

	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	request = &model.PaginationRequest{
		Limit: limit,
		Page:  page,
		Sort:  c.Query("sort"),
	}

	response, err = handler.contactService.GetContactList(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Contacts retrieved",
		"data":    response,
	})
}
