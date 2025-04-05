package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type JabatanDesaController struct {
	service *services.JabatanDesaService
}

func NewControllerJabatanDesa() *JabatanDesaController {
	return &JabatanDesaController{
		service: services.NewJabatanDesaService(),
	}
}
func (c *JabatanDesaController) GetData(ctx *gin.Context) {
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
	data, err := c.service.GetData(offset, limit)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	message := "Data kosong"
	if len(data) > 0 {
		message = "data berhasil dimuat"
	}
	response := resources.GetResponseJabatanDesa(data)
	resources.Success(ctx, message, response)

}

func (c *JabatanDesaController) CreateData(ctx *gin.Context) {
	var req requests.JabatanDesaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	err := c.service.CreateData(&req)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "data berhasil dimuat", req)
}
func (c *JabatanDesaController) UpdateData(ctx *gin.Context) {

	var req requests.JabatanDesaRequest
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

	err = c.service.UpdateData(&req, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			resources.NotFound(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "data berhasil diupdate", req)
}
func (c *JabatanDesaController) DeleteData(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resources.BadRequest(ctx, err)
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
	resources.Success(ctx, "data berhasil dihapus")
}
