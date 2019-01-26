package controllers

import (
	"lapanganku/models"

	"github.com/gin-gonic/gin"
)

// AllVenue controller
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
			"message": "Venue wasn't found",
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
			"message": "Sorry something wrong with the database",
			"error":   err,
		})
	} else {
		c.JSON(200, gin.H{
			"data": venue,
		})
	}
}

// UpdateVenue controller
func UpdateVenue(c *gin.Context) {

	var venue models.Venue

	err := models.ShowVenue(&venue, c.Params.ByName("id"))

	// Handle not found error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Venue not found",
		})
	}

	// c.BindJSON(&venue)

	test := c.BindJSON(&venue)

	c.JSON(200, gin.H{
		"data": test,
	})

	// err = models.UpdateVenue(&venue, c.Params.ByName("id"))

	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"message": "Something went wrong",
	// 	})
	// } else {
	// 	c.JSON(200, gin.H{
	// 		"data":    &venue,
	// 		"message": "Update here",
	// 	})
}
