package service

import (
	"errors"
	"ticket-app/entity"
	"ticket-app/repository"
)

type CatagoryService struct {
	Repository repository.CatagoryRepository
}

func (service CatagoryService) Get(Id string) (*entity.Catagory, error) {

	catagory := service.Repository.FindById(Id)

	if catagory == nil {
		return nil, errors.New("kategori gak ketemu bro")
	} else {

		return catagory, nil
	}

}
