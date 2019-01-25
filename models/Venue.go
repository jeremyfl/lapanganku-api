package models

import "lapanganku/config"

var db = config.Database()

// GetAllVenue handler for show all Venue
func GetAllVenue(v *[]Venue) (err error) {
	if err = db.Find(v).Error; err != nil {
		return err
	}

	return nil
}

// ShowVenue return venue based by id
func ShowVenue(v *Venue, id string) (err error) {
	if err := db.First(v, id).Error; err != nil {
		return err
	}

	return nil
}

// StoreVenue give all
func StoreVenue(v *Venue) (err error) {
	if err = db.Create(v).Error; err != nil {
		return err
	}

	return nil
}

// UpdateVenue to update all venue based by id
func UpdateVenue(v *Venue) (err error) {
	if err = db.Update(v, 1).Error; err != nil {
		return err
	}

	return nil
}
