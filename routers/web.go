package routers

import (
	"apidanadesa/app/controllers"
	"apidanadesa/app/middleware"
	"apidanadesa/app/models"
	"apidanadesa/config"
	"github.com/gin-gonic/gin"
)

func setUpRouterPing(router *gin.RouterGroup) {

	welcome := router.Group("/welcome")
	welcome.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to apidanadesa",
		})
	})
	welcome.GET("/testdb", func(c *gin.Context) {
		db := config.GetDB()
		if db == nil {
			c.JSON(500, gin.H{
				"error": "Database tidak terkoneksi",
			})
			return
		}

		var result []models.Bidang
		err := db.Offset(0).Limit(10).Find(&result).Error
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": result,
		})
	})

}

// SetUpRouterBidang sets up the routing group for the "Bidang" master data.
// It includes endpoints for retrieving, creating, updating, deleting, and importing data from a CSV file.

func SetUpRouterBidang(api *gin.RouterGroup) {
	c := controllers.NewControllerBidang()
	bidang := api.Group("bidang")
	bidang.Use(middleware.JWTMiddleware())
	bidang.GET("", c.GetAllBidangs)
	bidang.POST("", c.SaveBidang)
	bidang.PUT("", c.UpdateBidang)
	bidang.DELETE("/", c.DeleteBidang)
	bidang.POST("/import-csv", c.ImportDataCsv)
}

func SetUpRouterSubBidang(e *gin.RouterGroup) {
	c := controllers.NewSubBidangController()
	group := e.Group("sub-bidang")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetSubBidangs)
	group.POST("", c.CreateSubBidang)
	group.PUT("/:id", c.UpdateSubBidang)
	group.DELETE("/:id", c.DeleSubBidang)
	group.POST("/import-csv", c.ImportDataCsv)
}
func SetUpRouterKegiatan(e *gin.RouterGroup) {
	c := controllers.NewKegiatanController()
	group := e.Group("kegiatan")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetKegiatans)
	group.POST("", c.CreateKegiatan)
	group.PUT("/:id", c.UpdateKegiatan)
	group.DELETE("/:id", c.DeleteKegiatan)
}
func SetUpRouterTahunAnggaran(e *gin.RouterGroup) {
	c := controllers.NewTahunAnggaranController()
	group := e.Group("tahun-anggaran")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetData)
	group.POST("", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
}
func SetUpRouterAuth(e *gin.RouterGroup) {
	c := controllers.NewAuthUserController()
	group := e.Group("auth")
	group.POST("/register", c.RegisterUser)
	group.POST("/login", c.Login)
}

// SetUpRouterJabatanDesa sets up the routing group for the "Jabatan Desa" (Village Position) master data.
// It includes endpoints for retrieving, creating, updating, deleting, and importing data from a CSV file.
func SetUpRouterJabatanDesa(e *gin.RouterGroup) {
	c := controllers.NewControllerJabatanDesa()
	group := e.Group("jabatan-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetData)
	group.POST("", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
	group.POST("/import-csv", c.ImportDataCsv)
}

// master data belanja

func SetUpRouterKelompokBelanjaDesa(e *gin.RouterGroup) {
	c := controllers.NewControllerKelompokBelanja()
	group := e.Group("kelompok-belanja-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetData)
	group.POST("", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
	group.POST("/import-csv", c.ImportDataCsv)
}
func SetUpRouterJenisBelanjaDesa(e *gin.RouterGroup) {
	c := controllers.NewControllerJenisBelanja()
	group := e.Group("jenis-belanja-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetAll)
	group.POST("", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
}

func SetUpRouterObjekBelanjaDesa(e *gin.RouterGroup) {
	c := controllers.NewControllerObjekBelanjaDesa()
	group := e.Group("objek-belanja-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetAll)
	group.POST("", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
}

// master data pendapatan

func SetUpRouterKelompokPendapatanDesa(e *gin.RouterGroup) {
	c := controllers.NewKelompokPendapatanController()
	group := e.Group("kelompok-pendapatan-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetData)
	group.POST("", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
	group.POST("/import-csv", c.ImportDataCsv)
}

func SetUpRouterJenisPendapatanDesa(e *gin.RouterGroup) {
	c := controllers.NewControllerJenisPendapatan()
	group := e.Group("jenis-pendapatan-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetAll)
	group.POST("", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
}

func SetUpRouterObjekPendapatan(e *gin.RouterGroup) {
	c := controllers.NewObjekPendapatanController()
	group := e.Group("objek-pendapatan-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetAll)
	group.POST("", c.CreateData)
	group.PUT("/:id", c.UpdateData)
	group.DELETE("/:id", c.DeleteData)
}

func SetUpRouterPerangkatDesa(e *gin.RouterGroup) {
	c := controllers.NewPerangkatDesaController()
	group := e.Group("perangkat-desa")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetAll)
	group.POST("", c.SubmitForm)
	group.POST("/update/:id", c.UpdatePerangkat)
	group.DELETE("/:id", c.DeletePerangkat)
}

func SetUpRouterPreviewFoto(e *gin.RouterGroup) {
	e.GET("/preview/:filename", controllers.NewPerangkatDesaController().PreviewFile)
}

func SetUpRouterPolaKegiatan(e *gin.RouterGroup) {
	c := controllers.NewPolaKegiatanController()
	group := e.Group("pola-kegiatan")
	group.Use(middleware.JWTMiddleware())
	group.GET("", c.GetAll)
	group.POST("", c.Create)
	group.PUT("/:id", c.Update)
	group.DELETE("/:id", c.Delete)
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
	SetUpRouterObjekBelanjaDesa(api)
	SetUpRouterKelompokPendapatanDesa(api)
	SetUpRouterJenisPendapatanDesa(api)
	SetUpRouterObjekPendapatan(api)
	setUpRouterPing(api)
	SetUpRouterPerangkatDesa(api)
	SetUpRouterPreviewFoto(api)
	SetUpRouterPolaKegiatan(api)
}
