package models

import "gopkg.in/mgo.v2/bson"

type Transaction struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Field bson.ObjectId `json:"field" bson:"field"`
	Price int           `json:"price" bson:"price"`
}