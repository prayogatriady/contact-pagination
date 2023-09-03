package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/handler/helpers"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/service"
)

type ContactHandler interface {
	GetContactList(c *gin.Context)
	GetContact(c *gin.Context)
	CreateContact(c *gin.Context)
	UpdateContact(c *gin.Context)
	DeleteContact(c *gin.Context)
}

type contactHandler struct {
	contactService service.ContactService
	validator      *validator.Validate
}

func NewContactHandler(contactService service.ContactService) ContactHandler {
	return &contactHandler{
		contactService: contactService,
		validator: validator.New(),
	}
}

func (handler *contactHandler) GetContactList(c *gin.Context) {
	var (
		requestQuery  *model.PaginationRequest
		response *model.PaginationResponse
		limit    int
		page     int
		sort	 string
		err      error
	)

	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.StatusBadRequest_Query("limit"))
		return
	}

	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.StatusBadRequest_Query("page"))
		return
	}

	sort = c.Query("sort")
	if ok := helpers.ValidateSort(sort); !ok {
		c.JSON(http.StatusBadRequest, helpers.StatusBadRequest_Query("sort"))
		return
	}

	// if err := c.ShouldBindQuery(&requestQuery); err != nil {
	// 	c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
	// 	return
	// }

	requestQuery = &model.PaginationRequest{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

	response, err = handler.contactService.GetContactList(requestQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		return
	}

	c.JSON(http.StatusOK, helpers.StatusOKWithData("Contacts retrieved", response))
}

func (handler *contactHandler) GetContact(c *gin.Context) {
	var (
		response *model.ContactResponse
		contactId int
		err      error
	)

	contactId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.StatusBadRequest_Param("id"))
		return
	}

	response, err = handler.contactService.GetContact(contactId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		return
	}

	c.JSON(http.StatusOK, helpers.StatusOKWithData("Contact retrieved", response))
}

func (handler *contactHandler) CreateContact(c *gin.Context) {
	var (
		request *model.CreateContactRequest
		err      error
	)

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		return
	}

	if err = handler.contactService.CreateContact(request); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		return
	}

	c.JSON(http.StatusOK, helpers.StatusOK("Contact created"))
}

func (handler *contactHandler) UpdateContact(c *gin.Context) {
	var (
		request *model.UpdateContactRequest
		contactId int
		err      error
	)

	contactId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.StatusBadRequest_Param("id"))
		return
	}

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		return
	}

	if err = handler.contactService.UpdateContact(contactId, request); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		return
	}

	c.JSON(http.StatusOK, helpers.StatusOK("Contact updated"))
}

func (handler *contactHandler) DeleteContact(c *gin.Context) {
	var (
		contactId int
		err      error
	)

	contactId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.StatusBadRequest_Param("id"))
		return
	}

	if err = handler.contactService.DeleteContact(contactId); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		return
	}

	c.JSON(http.StatusOK, helpers.StatusOK("Contact deleted"))
}