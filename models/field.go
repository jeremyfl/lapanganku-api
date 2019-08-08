package models

import (
	"github.com/jeremylombogia/lapanganku-api/config"
	"gopkg.in/mgo.v2/bson"
)

type Field struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Name         string        `bson:"name" json:"name"`
	Venue        bson.ObjectId `json:"venue" bson:"venue"`
	Availability bool          `bson:"avaibility" json:"avaibility"`
	Price        int           `bson:"price" json:"price"`
}

const DOCUMENT string = "field"

var collection = session.DB(config.COLLECTION).C(DOCUMENT)

func StoreField(field *Field) error {
	return collection.Insert(&field)
}

func FetchField() ([]Field, error) {
	var field []Field

	err := collection.Find(bson.M{}).All(&field)

	return field, err
}

func FindField(id int) (Field, error) {
	var field Field

	// TODO:: Remove this implementation with dynamic payload
	err := collection.Find(bson.M{"id": id}).One(&field)

	return field, err
}
