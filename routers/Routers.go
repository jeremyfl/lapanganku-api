package routers

import (
	"lapanganku/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter main init router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("lapangan", controllers.AllVenue)
		v1.GET("lapangan/:id", controllers.ShowVenue)
		v1.POST("lapangan", controllers.AddVenue)
		v1.PATCH("lapangan/:id", controllers.UpdateVenue)
	}

	return r
}
