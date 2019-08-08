package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jeremylombogia/lapanganku-api/partner"
)

// SetupRouter main init router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("lapangan", controllers.AllVenue)
	// 	v1.GET("lapangan/:id", controllers.ShowVenue)
	// 	v1.POST("lapangan", controllers.AddVenue)
	// 	v1.PATCH("lapangan/:id", controllers.UpdateVenue)
	// }

	p := r.Group("/partner")
	{
		p.GET("transaction", partner.IndexTransaction)
		p.GET("transaction/:id", partner.ShowTransaction)
		p.POST("transaction", partner.StoreTransaction)
	}

	return r
}
