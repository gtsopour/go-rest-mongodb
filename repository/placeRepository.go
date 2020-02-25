package repository

import (
	"fmt"
	"go-rest-mongodb/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type placeRepository struct {
	db *mgo.Database
}

const Collection = "Places"

func PlaceRepository(db *mgo.Database) *placeRepository {
	return &placeRepository {
		db: db,
	}
}

func (p *placeRepository) FindAll() ([]models.Place, error) {
	var places []models.Place
	err := p.db.C(Collection).Find(bson.M{}).All(&places)
	return places, err
}

func (p *placeRepository) Insert(place models.Place) error {
	fmt.Println(place)
	err := p.db.C(Collection).Insert(&place)
	return err
}


