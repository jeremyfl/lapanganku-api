package main

import (
	"lapanganku/Routers"
	"lapanganku/config"
	"lapanganku/models"
)

var err error

func main() {

	config.Database().Close()
	config.Database().AutoMigrate(&models.Venue{})

	r := routers.SetupRouter()
	r.Run()
}
