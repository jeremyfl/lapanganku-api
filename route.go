package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeremylombogia/lapanganku-api/partner"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	partnerRoute := r.Group("/partner")
	{
		partnerRoute.GET("transaction", partner.IndexTransaction)
		partnerRoute.GET("transaction/:id", partner.ShowTransaction)
		partnerRoute.POST("transaction", partner.StoreTransaction)
	}

	return r
}
