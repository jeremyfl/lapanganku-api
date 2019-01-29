package main

import (
	"lapanganku/config"
	"lapanganku/models"
	"lapanganku/routers"
)

var err error

func main() {

	config.Database().Close()
	config.Database().AutoMigrate(&models.Venue{})

	r := routers.SetupRouter()
	r.Run()
}
