package healthStatus

import (
	"github.com/gin-gonic/gin"
)

type healthStatus struct{}

/*
Ping return a message pong
*/
func Ping(c *gin.Context) {
	h := healthStatus{}
	h.ping(c)
}

func (h healthStatus) ping(c *gin.Context) {
	//return message "pong" and status code 200
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
