package routers

import (
	"apidanadesa/app/controllers"
	"apidanadesa/app/middleware"
	"github.com/gin-gonic/gin"
)

func setUpRouterPing(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bidang",
		})
	})

	router.GET("/bidangs", controllers.NewControllerBidang().GetAllBidangs)

}
func SetUpRouterBidang(api *gin.RouterGroup) {
	c := controllers.NewControllerBidang()
	bidang := api.Group("/bidang")
	bidang.Use(middleware.JWTMiddleware())
	bidang.GET("/", c.GetAllBidangs)
	bidang.POST("/", c.SaveBidang)
	bidang.PUT("/", c.UpdateBidang)
	bidang.DELETE("/", c.DeleteBidang)
	bidang.POST("/import-csv", c.ImportDataCsv)
}

func SetUpRouterSubBidang(e *gin.RouterGroup) {
	c := controllers.NewSubBidangController()
	group := e.Group("/sub-bidang")
	group.Use(middleware.JWTMiddleware())
	group.GET("/", c.GetSubBidangs)
	group.POST("/", c.CreateSubBidang)
	group.PUT("/:id", c.UpdateSubBidang)
	group.DELETE("/:id", c.DeleSubBidang)
	group.POST("/import-csv", c.ImportDataCsv)
}
func SetUpRouterKegiatan(e *gin.RouterGroup) {
	c := controllers.NewKegiatanController()
	group := e.Group("/kegiatan")
	group.Use(middleware.JWTMiddleware())
	group.GET("/", c.GetKegiatans)
	group.POST("/", c.CreateKegiatan)
	group.PUT("/:id", c.UpdateKegiatan)
	group.DELETE("/:id", c.DeleteKegiatan)
}
func SetUpRouterTahunAnggaran(e *gin.RouterGroup) {
	c := controllers.NewTahunAnggaranController()
	group := e.Group("/tahun-anggaran")
	group.Use(middleware.JWTMiddleware())
	group.GET("/", c.GetData)
	group.POST("/", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
}
func SetUpRouterAuth(e *gin.RouterGroup) {
	c := controllers.NewAuthUserController()
	group := e.Group("/auth")
	group.POST("/register", c.RegisterUser)
	group.POST("/login", c.Login)
}
func SetUpRouterJabatanDesa(e *gin.RouterGroup) {
	c := controllers.NewControllerJabatanDesa()
	group := e.Group("/jabatan-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("/", c.GetData)
	group.POST("/", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
	group.POST("/import-csv", c.ImportDataCsv)
}
func SetUpRouterKelompokBelanjaDesa(e *gin.RouterGroup) {
	c := controllers.NewControllerKelompokBelanja()
	group := e.Group("/kelompok-belanja-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("/", c.GetData)
	group.POST("/", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
	group.POST("/import-csv", c.ImportDataCsv)
}
func SetUpRouterJenisBelanjaDesa(e *gin.RouterGroup) {
	c := controllers.NewControllerJenisBelanja()
	group := e.Group("/jenis-belanja-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("/", c.GetAll)
	group.POST("/", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
}
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	SetUpRouterBidang(api)
	SetUpRouterSubBidang(api)
	SetUpRouterKegiatan(api)
	SetUpRouterTahunAnggaran(api)
	SetUpRouterAuth(api)
	SetUpRouterJabatanDesa(api)
	SetUpRouterKelompokBelanjaDesa(api)
	SetUpRouterJenisBelanjaDesa(api)
}
