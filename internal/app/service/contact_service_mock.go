package service

import (
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"github.com/stretchr/testify/mock"
)

type ContactServiceMock struct {
	mock.Mock
}

func (m *ContactServiceMock) Paginate(request *model.PaginationRequest) (*model.PaginationResponse, error) {
	args := m.Called(request)
	if args.Get(0) != nil {
		return args.Get(0).(*model.PaginationResponse), args.Error(1)
	}
	return nil, args.Error(1)
}
