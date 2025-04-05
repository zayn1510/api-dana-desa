package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type SubBidangController struct {
	service *services.SubBidangService
}

func NewSubBidangController() *SubBidangController {
	return &SubBidangController{
		service: services.NewSubBidangService(),
	}
}
func (c *SubBidangController) GetSubBidangs(ctx *gin.Context) {
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
	data, err := c.service.GetAllSubBidang(offset, limit)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	message := "Data kosong"
	if len(data) > 0 {
		message = "data berhasil dimuat"
	}
	response := resources.GetSubBidangResponse(data)
	resources.Success(ctx, message, response)

}

func (c *SubBidangController) CreateSubBidang(ctx *gin.Context) {
	var req requests.SubBidangRequestCreate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	err := c.service.CreateSubBidang(&req)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "data berhasil dibuat", req)
}
func (c *SubBidangController) UpdateSubBidang(ctx *gin.Context) {

	var req requests.SubBidangRequestCreate
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
	err = c.service.UpdateSubBidang(&req, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "sudah digunakan") {
			resources.Conflict(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "data berhasil diupdate", req)
}
func (c *SubBidangController) DeleSubBidang(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	err = c.service.DeleteSubBidang(uint(id))
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
