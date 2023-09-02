package repository

import (
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/stretchr/testify/mock"
)

type ContactRepositoryMock struct {
	mock.Mock
}

func (m *ContactRepositoryMock) GetContactList(pagination *model.Pagination) ([]*model.Contact, int64, error) {
    args := m.Called(pagination)
    return args.Get(0).([]*model.Contact), args.Get(1).(int64), args.Error(2)

}

func (m *ContactRepositoryMock) GetContact(contactId int) (*model.Contact, error) {
	args := m.Called(contactId)
    return args.Get(0).(*model.Contact), args.Error(1)

}

func (m *ContactRepositoryMock) CreateContact(contact *model.Contact) (err error) {
    args := m.Called(contact)
    return args.Error(0)
}

func (m *ContactRepositoryMock) UpdateContact(contact *model.Contact) (err error) {
	args := m.Called(contact)
    return args.Error(0)
}

func (m *ContactRepositoryMock) DeleteContact(contactId int) (err error) {
	args := m.Called(contactId)
    return args.Error(0)
}