package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Venue struct {
	ID           bson.ObjectId   `bson:"_id" json:"id"`
	Name         string          `bson:"name" json:"name"`
	Field        []bson.ObjectId `json:"field" bson:"field"`
	Address      string          `json:"address" bson:"address"`
	User         bson.ObjectId   `json:"user" bson:"user"`
	City         bson.ObjectId   `json:"city" bson:"city"`
	Availability bool            `bson:"avaibility" json:"avaibility"`
}
