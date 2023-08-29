package repository

import (
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/stretchr/testify/mock"
)

type ContactRepositoryMock struct {
	mock.Mock
}

func (r *ContactRepositoryMock) Paginate(offset int, limit int, sort string) ([]*model.Contact, error) {
	arguments := r.Called(offset, limit, sort)

	var (
		return1 []*model.Contact
		return2 error
	)

	if arguments.Get(0) != nil {
		return1 = arguments.Get(0).([]*model.Contact)
	}
	if arguments.Get(1) != nil {
		return2 = arguments.Get(1).(error)
	}

	return return1, return2
}

func (r *ContactRepositoryMock) Count(value interface{}) (int64, error) {
	arguments := r.Called(value)

	var (
		return1 int64
		return2 error
	)

	if arguments.Get(0) != nil {
		return1 = arguments.Get(0).(int64)
	}
	if arguments.Get(1) != nil {
		return2 = arguments.Get(1).(error)
	}

	return return1, return2
}
