package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root:dev@/lapanganku?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Can't connect to database", err)
	}

	return db
}

// var DB *gorm.DB
