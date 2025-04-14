package main

import (
	"apidanadesa/app/middleware"
	"apidanadesa/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	middleware.SetCors(router)
	routers.RegisterRoutes(router)
	router.Run(":8080")
}
