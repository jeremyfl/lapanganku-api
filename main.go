package main

import (
	"lapanganku/Routers"
	"lapanganku/config"
	"lapanganku/models"

	_ "github.com/go-sql-driver/mysql"
)

var err error

func main() {
	config.Database().Close()
	config.Database().AutoMigrate(&models.Venue{})

	r := routers.SetupRouter()
	r.Run()
}
