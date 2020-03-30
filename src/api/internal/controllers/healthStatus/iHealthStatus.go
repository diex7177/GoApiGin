package healthStatus

import "github.com/gin-gonic/gin"

// IHealthStatus interface
type IHealthStatus interface {
	Ping(c *gin.Context)
}
