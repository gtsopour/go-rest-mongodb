package repository

import (
// 	"log"
  "fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"go-rest-mongodb/model"
)

type placeRepository struct {
	db *mgo.Database
	collection string
}

func PlaceRepository(db *mgo.Database, collection string) *placeRepository {
	return &placeRepository {
		db: db,
		collection: collection,
	}
}

func (p *placeRepository) FindAll() ([]model.Place, error) {
	var places []model.Place
	err := p.db.C(p.collection).Find(bson.M{}).All(&places)
	fmt.Println(err)
	return places, err
}

func (p *placeRepository) Insert(place model.Place) error {
  fmt.Println(place)
	err := p.db.C(p.collection).Insert(&place)
	fmt.Println(err)
	return err
}


