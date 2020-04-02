package healthStatus

import (
	"github.com/gin-gonic/gin"
)

//Controller ...
type Controller struct {
	IHealthStatus
}

//SendPong ...
func (controller Controller) SendPong(c *gin.Context) {
	//return message "pong" and status code 200
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
