package service

import (
	"math"

	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/repository"
)

type ContactService interface {
	Paginate(request *model.PaginationRequest) (*model.PaginationResponse, error)
}

type contactService struct {
	contactRepository repository.ContactRepository
}

func NewContactService(contactRepository repository.ContactRepository) ContactService {
	return &contactService{
		contactRepository: contactRepository,
	}
}

func (service *contactService) Paginate(request *model.PaginationRequest) (*model.PaginationResponse, error) {
	var (
		pagination       *model.Pagination
		contacts         []*model.Contact
		contactResponse  *model.ContactResponse
		contactsResponse []*model.ContactResponse
		err              error
		totalRows        int64
		totalPages       int
		response         *model.PaginationResponse
	)

	pagination = &model.Pagination{
		Limit: request.Limit,
		Page:  request.Page,
		Sort:  request.Sort,
	}

	contacts, err = service.contactRepository.Paginate(pagination.GetOffset(), pagination.GetLimit(), pagination.GetSort())
	if err != nil {
		return nil, err
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

	totalRows, err = service.contactRepository.Count(contacts)
	if err != nil {
		return nil, err
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

	return response, nil
}
