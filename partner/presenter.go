package partner

import (
	"github.com/gin-gonic/gin"
)

type Presenter struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, p *Presenter) {
	c.JSON(p.Status, p)
}
