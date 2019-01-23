package models

import (
	"github.com/jinzhu/gorm"
)

// Venue basic model current structur of venue
type Venue struct {
	gorm.Model
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Availability bool   `json:"avaibility"`
}

// TableName for venue
func (b *Venue) TableName() string {
	return "venue"
}
