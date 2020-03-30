package main

import (
	"GoApiGin/src/api/internal/controllers/healthStatus"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", healthStatus.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
