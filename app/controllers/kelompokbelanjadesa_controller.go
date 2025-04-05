package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type KelompokBelanjaDesaController struct {
	service *services.KelompokBelanjaDesaService
}

func NewControllerKelompokBelanja() *KelompokBelanjaDesaController {
	return &KelompokBelanjaDesaController{
		service: services.NewKelompokBelanjaService(),
	}
}
func (c *KelompokBelanjaDesaController) GetData(ctx *gin.Context) {
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
	data, err := c.service.GetAll(offset, limit)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	message := "Data kosong"
	if len(data) > 0 {
		message = "data berhasil dimuat"
	}
	response := resources.GetResponseKelompokBelanja(data)
	resources.Success(ctx, message, response)

}

func (c *KelompokBelanjaDesaController) CreateData(ctx *gin.Context) {
	var req requests.KelompokBelanjaDesaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	if _, err := requests.Validate(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	err := c.service.CreateData(&req)
	if err != nil {
		if strings.Contains(err.Error(), "sudah digunakan") {
			resources.Conflict(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Created(ctx, "data berhasil dibuat", req)
}
func (c *KelompokBelanjaDesaController) UpdateData(ctx *gin.Context) {

	var req requests.KelompokBelanjaDesaRequest
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
		if strings.Contains(err.Error(), "sudah digunakan") {
			resources.Conflict(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "data berhasil diupdate", req)
}
func (c *KelompokBelanjaDesaController) DeleteData(ctx *gin.Context) {
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
	resources.Success(ctx, "data berhasil dibuat")
}
