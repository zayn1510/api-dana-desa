package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type JenisPendapatanController struct {
	service *services.JenisPendapatanService
}

func NewControllerJenisPendapatan() *JenisPendapatanController {
	return &JenisPendapatanController{
		service: services.NewJenisPendapatanService(),
	}
}

func (c *JenisPendapatanController) GetAll(ctx *gin.Context) {
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
		resources.InternalError(ctx, err)
		return
	}
	message := "data kosong"
	if len(data) > 0 {
		message = "data berhasil dimuat"
	}
	response := resources.GetJenisPendapatanResource(data)
	resources.Success(ctx, message, response)
}

func (c *JenisPendapatanController) CreateData(ctx *gin.Context) {
	var req = new(requests.JenisPendapatanRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	if err, validationErrors := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validationErrors)
		return
	}
	if err := c.service.Create(req); err != nil {
		if strings.Contains(err.Error(), "sudah digunakan") {
			resources.Conflict(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "data berhasil dibuat")
}

func (c *JenisPendapatanController) UpdateData(ctx *gin.Context) {
	req := new(requests.JenisPendapatanRequest)
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		requests.HandleBindError(ctx, err)
		return
	}
	if err, validationErrors := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validationErrors)
		return
	}
	if err := c.service.Update(uint(id), req); err != nil {
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
	resources.Success(ctx, "data berhasil di update", req)
}

func (c *JenisPendapatanController) DeleteData(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resources.BadRequest(ctx, err)
	}
	if err := c.service.Delete(uint(id)); err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			resources.NotFound(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
	}
	resources.Success(ctx, "data berhasil dihapus")
}
