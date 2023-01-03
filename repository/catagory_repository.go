package repository

import (
	"github.com/chandrafebrian/golang-real-chandra/entity"
)

type CatagoryRepository interface {
	FindById(id string) *entity.Catagory
}
