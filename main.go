package main

import (
	"github.com/jeremylombogia/lapanganku-api/config"
	"github.com/jeremylombogia/lapanganku-api/models"
	"github.com/jeremylombogia/lapanganku-api/routers"
)

var err error

func main() {
	config.Database().Close()
	config.Database().AutoMigrate(&models.Venue{}, &models.Transaction{})
	config.Database().Model(&models.Transaction{}).AddForeignKey("venue_id", "venues(id)", "cascade", "cascade")

	r := routers.SetupRouter()
	r.Run()
}
