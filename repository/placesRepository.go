package repository

import (
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

// Get all Places
func (p *placeRepository) FindAll() ([]models.Place, error) {
	var places []models.Place
	err := p.db.C(Collection).Find(bson.M{}).All(&places)
	return places, err
}

// Create a new Place
func (p *placeRepository) Insert(place models.Place) error {
	err := p.db.C(Collection).Insert(&place)
	return err
}

// Delete an existing Place
func (p *placeRepository) Delete(id string) error {
	err := p.db.C(Collection).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
