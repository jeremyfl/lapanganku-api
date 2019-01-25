package controllers

import (
	"lapanganku/models"

	"github.com/gin-gonic/gin"
)

// AllVenue list of venue available
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

// ShowVenue controller
func ShowVenue(c *gin.Context) {
	var venue models.Venue
	foundVenue := models.ShowVenue(&venue, c.Param("id"))

	if foundVenue != nil {
		c.JSON(404, gin.H{
			"message": "Book wasn't found",
		})
	} else {
		c.JSON(200, gin.H{
			"data": venue,
		})
	}
}

// AddVenue controller
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

// func UpdateVenue(c *gin.Context) {
// 	var venue models.Venue

// 	// c.JSON(500, gin.H{
// 	// 	"venue":   &venue,
// 	// 	"message": "Sorry something went wrong",
// 	// })

// 	// c.BindJSON(&venue)

// 	// err := models.UpdateVenue(&venue)

// 	// if err != nil {
// 	// 	c.JSON(200, gin.H{
// 	// 		"data": venue,
// 	// 	})
// 	// } else {
// 	// 	c.JSON(500, gin.H{
// 	// 		"message": "Sorry something went wrong",
// 	// 	})
// 	// }
// }
