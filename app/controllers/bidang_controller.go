package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	message := "data berhasil dimuat"
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
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: "Invalid request format",
			Status:  false,
		})
		return
	}
	err = c.service.CreateData(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Message: "data berhasil dibuat",
	})
}

func (c *BidangController) UpdateBidang(ctx *gin.Context) {
	var req requests.BidangRequestCreate
	idStr := ctx.DefaultQuery("id", "0")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: "Invalid request format",
			Status:  false,
		})
		return
	}

	err = c.service.UpdateData(&req, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Message: "data berhasil diupate",
	})
}
func (c *BidangController) DeleteBidang(ctx *gin.Context) {
	idStr := ctx.DefaultQuery("id", "0")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	err = c.service.DeleteData(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Message: "data berhasil dihapus",
	})
}
