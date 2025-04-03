package routers

import (
	"apidanadesa/app/controllers"
	"apidanadesa/app/middleware"
	"github.com/gin-gonic/gin"
)

type RouteSetupFunc func(*gin.Engine)

var routeSetups = []RouteSetupFunc{
	SetUpRouterBidang,
	SetUpRouterAuth,
	SetUpRouterSubBidang,
}

func SetUpRouterBidang(e *gin.Engine) {
	c := controllers.NewControllerBidang()
	group := e.Group("/bidang")
	group.Use(middleware.JWTMiddleware())
	group.GET("/", c.GetAllBidangs)
	group.POST("/", c.SaveBidang)
	group.PUT("/", c.UpdateBidang)
	group.DELETE("/", c.DeleteBidang)
}

func SetUpRouterSubBidang(e *gin.Engine) {
	c := controllers.NewSubBidangController()
	group := e.Group("/sub-bidang")
	group.Use(middleware.JWTMiddleware())
	group.GET("/", c.GetSubBidangs)
	group.POST("/", c.CreateSubBidang)
	group.PUT("/:id", c.UpdateSubBidang)
	group.DELETE("/:id", c.DeleSubBidang)
}

func SetUpRouterAuth(e *gin.Engine) {
	c := controllers.NewAuthUserController()
	group := e.Group("/auth")
	group.POST("/register", c.RegisterUser)
	group.POST("/login", c.Login)

}
func RegisterRoutes(r *gin.Engine) {
	for _, setup := range routeSetups {
		setup(r)
	}
}
