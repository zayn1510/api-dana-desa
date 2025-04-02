package routers

import (
	"apidanadesa/app/controllers"
	"github.com/gin-gonic/gin"
)

type RouteSetupFunc func(*gin.Engine)

var routeSetups = []RouteSetupFunc{
	SetUpRouterBidang,
}

func SetUpRouterBidang(e *gin.Engine) {

	c := controllers.NewControllerBidang()
	group := e.Group("/bidang")
	group.GET("/", c.GetAllBidangs)
	group.POST("/", c.SaveBidang)
	group.PUT("/", c.UpdateBidang)
	group.DELETE("/", c.DeleteBidang)
}
func RegisterRoutes(r *gin.Engine) {
	for _, setup := range routeSetups {
		setup(r)
	}
}
