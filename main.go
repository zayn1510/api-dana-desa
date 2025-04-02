package main

import (
	"apidanadesa/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	c.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongs",
		})
	})
	routers.RegisterRoutes(c)
	c.Run(":8080")
}
