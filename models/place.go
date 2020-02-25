package models

import "gopkg.in/mgo.v2/bson"

type Place struct {
	ID          bson.ObjectId	`bson:"_id" json:"id"`
	Title       string			`bson:"title" json:"title"`
	Description string			`bson:"description" json:"description"`
}
