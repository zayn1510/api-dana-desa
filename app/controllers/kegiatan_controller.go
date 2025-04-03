package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type KegiatanController struct {
	service *services.KegiatanService
}

func NewKegiatanController() *KegiatanController {
	return &KegiatanController{
		service: services.NewKegiatanService(),
	}
}
func (c *KegiatanController) GetKegiatans(ctx *gin.Context) {
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
	data, err := c.service.GetAllKegiatan(offset, limit)
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
	response := resources.GetKegiatanResponse(data)
	ctx.JSON(http.StatusOK, resources.Response{
		Message: message,
		Status:  true,
		Data:    response,
	})

}

func (c *KegiatanController) CreateKegiatan(ctx *gin.Context) {
	var req requests.KegiatanRequestCreate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	err := c.service.CreateKegiatan(&req)
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
func (c *KegiatanController) UpdateKegiatan(ctx *gin.Context) {

	var req requests.KegiatanRequestCreate
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

	err = c.service.UpdateKegiatan(&req, uint(id))
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
func (c *KegiatanController) DeleteKegiatan(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	err = c.service.DeleteKegiatan(uint(id))
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
