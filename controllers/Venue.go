package controllers

import (
	"lapanganku/models"

	"github.com/gin-gonic/gin"
)

func AllVenue(c *gin.Context) {
	var venue []models.Venue

	err := models.GetAllVenue(&venue)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "Venue not found",
		})
	} else {
		c.JSON(200, gin.H{
			"data": venue,
		})

	}
}

func AddVenue(c *gin.Context) {
	var venue models.Venue

	c.BindJSON(&venue)

	err := models.StoreVenue(&venue)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Sorry something wrong",
		})
	} else {
		c.JSON(200, gin.H{
			"data": venue,
		})
	}
}
