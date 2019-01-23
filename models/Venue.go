package models

import "lapanganku/config"

var db = config.Database()

func GetAllVenue(v *[]Venue) (err error) {
	if err = db.Find(v).Error; err != nil {
		return err
	}

	return nil
}

func StoreVenue(v *Venue) (err error) {
	if err = db.Create(v).Error; err != nil {
		return err
	}

	return nil
}
