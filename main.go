package main

import (
	"fmt"
	"lapanganku/Routers"
	"lapanganku/config"
	"lapanganku/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	config.DB, err = gorm.Open("mysql", "root:dev@/lapanganku?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Venue{})

	r := routers.SetupRouter()
	r.Run()
}
