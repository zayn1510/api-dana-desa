package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	offset := (page - 1) * limit
	data, err := c.service.GetAllSubBidang(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	message := "Data kosong"
	if len(data) > 0 {
		message = "data berhasil dimuat"
	}
	response := resources.GetSubBidangResponse(data)
	ctx.JSON(http.StatusOK, resources.Response{
		Message: message,
		Status:  true,
		Data:    response,
	})

}

func (c *SubBidangController) CreateSubBidang(ctx *gin.Context) {
	var req requests.SubBidangRequestCreate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	err := c.service.CreateSubBidang(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Data:    req,
		Message: "data berhasil dibuat",
	})
}
func (c *SubBidangController) UpdateSubBidang(ctx *gin.Context) {

	var req requests.SubBidangRequestCreate
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}

	err = c.service.UpdateSubBidang(&req, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Data:    req,
		Message: "data berhasil diperbarui",
	})
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
