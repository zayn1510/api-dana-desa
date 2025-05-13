package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type AnggaranKegiatanController struct {
	service *services.AnggaranKegiatanService
}

func NewAnggaranKegiatanController() *AnggaranKegiatanController {
	return &AnggaranKegiatanController{
		service: services.NewAnggaranKegiatanService(),
	}
}
func (c *AnggaranKegiatanController) GetAll(ctx *gin.Context) {

	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	// Calculate offset for pagination
	offset := (page - 1) * limit
	data, err := c.service.GetData(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	message := "data berhasil dimuat"
	if len(data) == 0 {
		message = "Data kosong"
	}
	bidang := resources.GetAnggaranKegiatanResource(data)
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Message: message,
		Data:    bidang,
	})
}

func (c *AnggaranKegiatanController) Save(ctx *gin.Context) {
	var req requests.AnggaranKegiatanRequestCreate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		requests.HandleBindError(ctx, err)
		return
	}

	if err, validationErrors := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validationErrors)
		return
	}
	if err := c.service.CreateData(&req); err != nil {
		if strings.Contains(err.Error(), "sudah digunakan") {
			resources.Conflict(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "anggaran kegiatan berhasil dibuat")
}

func (c *AnggaranKegiatanController) Update(ctx *gin.Context) {
	var req requests.AnggaranKegiatanRequestUpdate
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	err = c.service.UpdateData(&req, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			resources.NotFound(ctx, err)
			return
		}
		if strings.Contains(err.Error(), "sudah digunakan") {
			resources.Conflict(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "anggaran kegiatan berhasil dibuat")
}
func (c *AnggaranKegiatanController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	err = c.service.DeleteData(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			resources.NotFound(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "anggaran kegiatan berhasil dihapus")
}
