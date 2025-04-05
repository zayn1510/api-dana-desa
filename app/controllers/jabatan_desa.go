package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	data, err := c.service.GetData(offset, limit)
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
	response := resources.GetResponseJabatanDesa(data)
	ctx.JSON(http.StatusOK, resources.Response{
		Message: message,
		Status:  true,
		Data:    response,
	})

}

func (c *JabatanDesaController) CreateData(ctx *gin.Context) {
	var req requests.JabatanDesaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	err := c.service.CreateData(&req)
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
func (c *JabatanDesaController) UpdateData(ctx *gin.Context) {

	var req requests.JabatanDesaRequest
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
		Data:    req,
		Message: "data berhasil diperbarui",
	})
}
func (c *JabatanDesaController) DeleteData(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
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
