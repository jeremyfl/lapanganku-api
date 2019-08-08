package models

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	VenueID uint   `gorm:"not null" json:"venueId"`
	Venue   Venue  `json:"venue"`
	User    string `json:"user" xml:"user" form:"user" query:"user"`
	Status  int    `json:"status" xml:"status" form:"status" query:"status"`
}

func Fetch(transactions *[]Transaction) (err error) {
	if err = db.Preload("Venue").Find(&transactions).Error; err != nil {
		return err
	}

	return nil
}

func Show(transaction *Transaction, id string) (err error) {
	if err = db.Preload("Venue").Where("id = ?", id).First(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func Create(transaction *Transaction) (err error) {
	if err = db.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}
