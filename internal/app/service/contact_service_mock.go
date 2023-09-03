package service

import (
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/stretchr/testify/mock"
)

type ContactServiceMock struct {
	mock.Mock
}

func (m *ContactServiceMock) GetContactList(request *model.PaginationRequest) (*model.PaginationResponse, error) {
	args := m.Called(request)
	if args.Get(0) != nil {
		return args.Get(0).(*model.PaginationResponse), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ContactServiceMock) GetContact(contactId int) (response *model.ContactResponse, err error) {
	args := m.Called(contactId)
	if args.Get(0) != nil {
		return args.Get(0).(*model.ContactResponse), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ContactServiceMock) CreateContact(request *model.ContactRequest) (err error) {
	args := m.Called(request)
	return args.Error(0)
}

func (m *ContactServiceMock) UpdateContact(contactId int, request *model.ContactRequest) (err error) {
	args := m.Called(contactId, request)
	return args.Error(0)
}

func (m *ContactServiceMock) DeleteContact(contactId int) (err error) {
	args := m.Called(contactId)
	return args.Error(0)
}