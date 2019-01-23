package models

import (
	"github.com/jinzhu/gorm"
)

type Venue struct {
	gorm.Model
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Availability bool   `json:"avaibility"`
}

func (b *Venue) TableName() string {
	return "venue"
}
