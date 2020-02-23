package models

import (
  "gopkg.in/mgo.v2/bson"
)
// Represents a place, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type place struct {
	id          bson.ObjectId `bson:"_id" json:"id"`
	name        string        `bson:"name" json:"name"`
	description string        `bson:"description" json:"description"`
}
