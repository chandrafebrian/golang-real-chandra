package repository

import (
	"github.com/chandrafebrian/golang-real-chandra/entity"

	"github.com/stretchr/testify/mock"
)

type CatagoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CatagoryRepositoryMock) FindById(id string) *entity.Catagory {

	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}

	catagory := arguments.Get(0).(entity.Catagory)
	return &catagory
}
