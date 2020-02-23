package repository

import (
  "go-rest-mongodb/model"
)

type IPlaceRepository interface {
	Insert(*model.Place) error
// 	Update(string, *model.Place) error
// 	Delete(string) error
// 	FindByID(string) (*model.Place, error)
	FindAll() ([]model.Place, error)
}
