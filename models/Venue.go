package models

import "lapanganku/config"

func GetAllVenue(v *[]Venue) (err error) {
	if err = config.DB.Find(v).Error; err != nil {
		return err
	}

	return nil
}

func StoreVenue(v *Venue) (err error) {
	if err = config.DB.Create(v).Error; err != nil {
		return err
	}
	return nil
}
