package models

import (
	"github.com/jinzhu/gorm"
)

type Venue struct {
	gorm.Model
	Name         string
	Price        int
	Availability bool
}

func (b *Venue) TableName() string {
	return "venue"
}
