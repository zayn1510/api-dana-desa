package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusBadRequest,
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusBadRequest,
		})
		return
	}
	offset := (page - 1) * limit
	data, err := c.service.GetAll(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusInternalServerError,
		})
		return
	}
	message := "Data kosong"
	if len(data) > 0 {
		message = "data berhasil dimuat"
	}
	response := resources.GetResponseKelompokBelanja(data)
	ctx.JSON(http.StatusOK, resources.Response{
		Message: message,
		Status:  true,
		Code:    http.StatusOK,
		Data:    response,
	})

}

func (c *KelompokBelanjaDesaController) CreateData(ctx *gin.Context) {
	var req requests.KelompokBelanjaDesaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusBadRequest,
		})
		return
	}

	if errors, err := requests.Validate(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"validation_errors": errors})
		return
	}
	if msg, exists := c.service.IsKodeExist(req.Kode, uint(0)); exists {
		ctx.JSON(http.StatusConflict, resources.Response{
			Message:   msg,
			Status:    false,
			Code:      http.StatusConflict,
			Duplicate: true,
		})
		return
	}
	err := c.service.CreateData(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Data:    req,
		Code:    http.StatusOK,
		Message: "data berhasil dibuat",
	})
}
func (c *KelompokBelanjaDesaController) UpdateData(ctx *gin.Context) {

	var req requests.KelompokBelanjaDesaRequest
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusBadRequest,
		})
		return
	}

	msg, exist := c.service.IsKodeExist(req.Kode, uint(id))
	if exist {
		ctx.JSON(http.StatusConflict, resources.Response{
			Message:   msg,
			Status:    false,
			Code:      http.StatusConflict,
			Duplicate: true,
		})
		return
	}
	err = c.service.UpdateData(&req, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Data:    req,
		Code:    http.StatusOK,
		Message: "data berhasil diperbarui",
	})
}
func (c *KelompokBelanjaDesaController) DeleteData(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusBadRequest,
		})
		return
	}
	err = c.service.DeleteData(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
			Code:    http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Message: "data berhasil dihapus",
		Code:    http.StatusOK,
	})
}
