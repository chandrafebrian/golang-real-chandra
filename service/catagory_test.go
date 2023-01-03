package service

import (
	"testing"

	"github.com/chandrafebrian/golang-real-chandra/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var catagoryRepository = &repository.CatagoryRepositoryMock{Mock: mock.Mock{}}
var catagoryService = CatagoryService{Repository: catagoryRepository}

func TestCatagoryService_GetErr(t *testing.T) {

	// program mock
	catagoryRepository.Mock.On("FindById", "1").Return(nil)
	catagory, err := catagoryService.Get("1")
	assert.Nil(t, catagory)
	assert.NotNil(t, err)
}

// func TestCatagoryService_Sucssess(t *testing.T) {

// 	catagory := entity.Catagory{
// 		Id:   "1",
// 		Name: "laptop",
// 	}

// 	catagoryRepository.Mock.On("FindById", "2").Return(nil)

// 	hasil, err := catagoryService.Get("2")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, hasil)
// 	assert.Equal(t, catagory.Id, hasil.Id)
// 	assert.Equal(t, catagory.Name, hasil.Name)
// }
