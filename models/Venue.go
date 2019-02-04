package models

import (
	"lapanganku/config"

	"github.com/jinzhu/gorm"
)

// Venue model
type Venue struct {
	gorm.Model
	Name         string `json:"name"`
	Address      string `json:"address"`
	Price        int    `json:"price"`
	Availability bool   `json:"avaibility"`
}

var db = config.Database()

// GetAllVenue model
func GetAllVenue(v *[]Venue) (err error) {
	if err = db.Find(v).Error; err != nil {
		return err
	}

	return nil
}

// ShowVenue model
func ShowVenue(v *Venue, id string) (err error) {
	if err := db.Where("id = ?", id).First(v).Error; err != nil {
		return err
	}

	return nil
}

// StoreVenue model
func StoreVenue(v *Venue) (err error) {
	if err = db.Create(v).Error; err != nil {
		return err
	}

	return nil
}

// UpdateVenue model
func UpdateVenue(v *Venue, id string) (err error) {
	if err = db.Save(v).Error; err != nil {
		return err
	}

	return nil
}
