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
	SetUpRouterKegiatan,
	SetUpRouterTahunAnggaran,
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
func SetUpRouterKegiatan(e *gin.Engine) {
	c := controllers.NewKegiatanController()
	group := e.Group("/kegiatan")
	group.Use(middleware.JWTMiddleware())
	group.GET("/", c.GetKegiatans)
	group.POST("/", c.CreateKegiatan)
	group.PUT("/:id", c.UpdateKegiatan)
	group.DELETE("/:id", c.DeleteKegiatan)
}
func SetUpRouterTahunAnggaran(e *gin.Engine) {
	c := controllers.NewTahunAnggaranController()
	group := e.Group("/tahun-anggaran")
	group.Use(middleware.JWTMiddleware())
	group.GET("/", c.GetData)
	group.POST("/", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
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
