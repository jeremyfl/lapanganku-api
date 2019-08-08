package partner

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jeremylombogia/lapanganku-api/models"
)

func IndexTransaction(c *gin.Context) {
	var transactions []models.Transaction

	if err := models.Fetch(&transactions); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	Response(c, &Presenter{
		Status:  200,
		Message: nil,
		Data:    transactions,
	})
}

func ShowTransaction(c *gin.Context) {
	var transaction models.Transaction

	if err := models.Show(&transaction, c.Param("id")); err != nil {
		fmt.Println(err)

		if transaction.ID == 0 {
			Response(c, &Presenter{
				Status:  404,
				Message: "Record not found",
				Data:    nil,
			})
		}

		return
	}

	Response(c, &Presenter{
		Status:  200,
		Message: nil,
		Data:    transaction,
	})

	return
}

func StoreTransaction(c *gin.Context) {
	var payload Payload

	c.BindJSON(&payload)

	var transaction = models.Transaction{
		VenueID: payload.Data.VenueID,
		User:    payload.Data.User,
	}

	if err := models.Create(&transaction); err != nil {
		os.Exit(3)
		c.JSON(500, gin.H{
			"err": err,
		})
	}

	Response(c, &Presenter{
		Status:  201,
		Message: "Transaction Created",
		Data:    transaction,
	})
}
