package mocks

import (
	"16_api_crud/entities"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) List() ([]*entities.Product, error) {
	args := r.Called()
	return args.Get(0).([]*entities.Product), args.Error(1)
}
