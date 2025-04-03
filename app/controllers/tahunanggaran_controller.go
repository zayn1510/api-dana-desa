package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	data, err := c.service.GetAllTahun(offset, limit)
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
	response := resources.GetTahunAnggaranResource(data)
	ctx.JSON(http.StatusOK, resources.Response{
		Message: message,
		Status:  true,
		Data:    response,
	})

}

func (c *TahunAnggaranController) CreateData(ctx *gin.Context) {
	var req requests.TahunAnggaranRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	err := c.service.CreateTahunAnggaran(&req)
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
func (c *TahunAnggaranController) UpdateData(ctx *gin.Context) {

	var req requests.TahunAnggaranRequest
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

	err = c.service.UpdateTahunAnggaran(&req, uint(id))
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
func (c *TahunAnggaranController) DeleteData(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	err = c.service.DeleteTahunAnggaran(uint(id))
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
