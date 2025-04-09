package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type ObjekBelanjaDesaController struct {
	service *services.ObjekBelanjaDesaService
}

func NewControllerObjekBelanjaDesa() *ObjekBelanjaDesaController {
	return &ObjekBelanjaDesaController{
		service: services.NewObjekBelanjaDesaService(),
	}
}

func (c *ObjekBelanjaDesaController) GetAll(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "10")
	pageStr := ctx.DefaultQuery("page", "1")

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
	response := resources.GetResponseObjekBelanjaDesa(data)
	resources.Success(ctx, "data berhasil dimuat", response)
}

func (c *ObjekBelanjaDesaController) CreateData(ctx *gin.Context) {
	var req = new(requests.ObjekBelanjaDesaRequest)
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

func (c *ObjekBelanjaDesaController) UpdateData(ctx *gin.Context) {
	req := new(requests.ObjekBelanjaDesaRequest)
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

func (c *ObjekBelanjaDesaController) DeleteData(ctx *gin.Context) {
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
