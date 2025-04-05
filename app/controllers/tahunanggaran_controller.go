package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type TahunAnggaranController struct {
	service *services.TahunAnggaranService
}

func NewTahunAnggaranController() *TahunAnggaranController {
	return &TahunAnggaranController{
		service: services.NewTahunAnggaranService(),
	}
}
func (c *TahunAnggaranController) GetData(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	offset := (page - 1) * limit
	data, err := c.service.GetAllTahun(offset, limit)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	message := "Data kosong"
	if len(data) > 0 {
		message = "data berhasil dimuat"
	}
	response := resources.GetTahunAnggaranResource(data)
	resources.Success(ctx, message, response)

}

func (c *TahunAnggaranController) CreateData(ctx *gin.Context) {
	var req requests.TahunAnggaranRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	err := c.service.CreateTahunAnggaran(&req)
	if err != nil {
		if strings.Contains(err.Error(), "sudah digunakan") {
			resources.Conflict(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "data berhasil dimuat", req)
}
func (c *TahunAnggaranController) UpdateData(ctx *gin.Context) {

	var req requests.TahunAnggaranRequest
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	err = c.service.UpdateTahunAnggaran(&req, uint(id))
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
	resources.Success(ctx, "data berhasil diupdate", req)
}
func (c *TahunAnggaranController) DeleteData(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	err = c.service.DeleteTahunAnggaran(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			resources.NotFound(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "data berhasil dihapus")
}
