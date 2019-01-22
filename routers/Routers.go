package routers

import (
	"lapanganku/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("lapangan", controllers.AllVenue)
		v1.POST("lapangan", controllers.AddVenue)
	}

	return r
}
