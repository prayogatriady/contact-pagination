package service

import (
	"math"

	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/repository"
)

type ContactService interface {
	GetContactList(request *model.PaginationRequest) (response *model.PaginationResponse, err error)
	GetContact(contactId int) (response *model.ContactResponse, err error)
	CreateContact(request *model.CreateContactRequest) (err error)
	UpdateContact(contactId int, request *model.UpdateContactRequest) (err error)
	DeleteContact(contactId int) (err error)
}

type contactService struct {
	contactRepository repository.ContactRepository
}

func NewContactService(contactRepository repository.ContactRepository) ContactService {
	return &contactService{
		contactRepository: contactRepository,
	}
}

func (service *contactService) GetContactList(request *model.PaginationRequest) (response *model.PaginationResponse, err error) {
	var (
		pagination       *model.Pagination
		contacts         []*model.Contact
		contactResponse  *model.ContactResponse
		contactsResponse []*model.ContactResponse
		totalRows        int64
		totalPages       int
	)

	pagination = &model.Pagination{
		Limit: request.Limit,
		Page:  request.Page,
		Sort:  request.Sort,
	}

	contacts, totalRows, err = service.contactRepository.GetContactList(pagination)
	if err != nil {
		return
	}

	for _, contact := range contacts {
		contactResponse = &model.ContactResponse{
			ID:        contact.ID,
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Email:     contact.Email,
			Phone:     contact.Phone,
		}
		contactsResponse = append(contactsResponse, contactResponse)
	}

	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	response = &model.PaginationResponse{
		Limit:      pagination.GetLimit(),
		Page:       pagination.GetPage(),
		Sort:       pagination.GetSort(),
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Rows:       contactsResponse,
	}

	return
}

func (service *contactService) GetContact(contactId int) (response *model.ContactResponse, err error) {

	var (
		contact *model.Contact
	)

	contact, err = service.contactRepository.GetContact(contactId)
	if err != nil {
		return
	}

	response = &model.ContactResponse{
		ID: contact.ID,
		FirstName: contact.FirstName,
		LastName: contact.LastName,
		Email: contact.Email,
		Phone: contact.Phone,
	}

	return
}

func (service *contactService) CreateContact(request *model.CreateContactRequest) (err error) {

	var (
		contact *model.Contact
	)

	contact = &model.Contact{
		FirstName: request.FirstName,
		LastName: request.LastName,
		Email: request.Email,
		Phone: request.Phone,
	}

	err = service.contactRepository.CreateContact(contact)

	return
}

func (service *contactService) UpdateContact(contactId int, request *model.UpdateContactRequest) (err error) {

	var (
		contact *model.Contact
	)

	contact = &model.Contact{
		ID: contactId,
		FirstName: request.FirstName,
		LastName: request.LastName,
		Email: request.Email,
		Phone: request.Phone,
	}

	err = service.contactRepository.UpdateContact(contact)
	
	return
}

func (service *contactService) DeleteContact(contactId int) (err error) {

	err = service.contactRepository.DeleteContact(contactId)
	return
}