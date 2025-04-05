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

type BidangController struct {
	service *services.BidangService
}

func NewControllerBidang() *BidangController {
	return &BidangController{
		service: services.NewBidangService(),
	}
}

func (c *BidangController) GetAllBidangs(ctx *gin.Context) {

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
	message := "data berhasil dimuats"
	if len(data) == 0 {
		message = "Data kosong"
	}
	bidang := resources.GetBidangResource(data)
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Message: message,
		Data:    bidang,
	})
}

func (c *BidangController) SaveBidang(ctx *gin.Context) {
	var req requests.BidangRequestCreate
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	err = c.service.CreateData(&req)
	if err != nil {
		if strings.Contains(err.Error(), "sudah digunakan") {
			resources.Conflict(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "bidang berhasil dibuat")
}

func (c *BidangController) UpdateBidang(ctx *gin.Context) {
	var req requests.BidangRequestCreate
	idStr := ctx.DefaultQuery("id", "0")
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
	resources.Success(ctx, "bidang berhasil dibuat")
}
func (c *BidangController) DeleteBidang(ctx *gin.Context) {
	idStr := ctx.DefaultQuery("id", "0")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	err = c.service.DeleteData(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			resources.NotFound(ctx, err)
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "bidang berhasil dihapus")
}
