package repository

import (
	"ticket-app/entity"
)

type CatagoryRepository interface {
	FindById(id string) *entity.Catagory
}
