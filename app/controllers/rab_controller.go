package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RabController struct {
	service *services.RabService
}

func NewControllerRab() *RabController {
	return &RabController{
		service: services.NewRabService(),
	}
}

func (c *RabController) GetAll(ctx *gin.Context) {
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
	offset := (page - 1) * limit
	data, err := c.service.GetData(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	message := "Data berhasil dimuat"
	if len(data) == 0 {
		message = "Data kosong"
	}
	rabResources := resources.GetRabResource(data)
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Message: message,
		Data:    rabResources,
	})
}

func (c *RabController) Create(ctx *gin.Context) {
	var request requests.RabRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: "Invalid request data",
			Status:  false,
		})
		return
	}
	if err,validationErr := requests.Validate(request);err != nil {
		resources.BadRequest(ctx, validationErr)
		return
	}
	rab := request.ToModel()
	if err := c.service.Create(rab); err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, resources.Response{
		Status:  true,
		Message: "RAB created successfully",
	})
}

func (c *RabController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: "Invalid ID format",
			Status:  false,
		})
		return
	}

	var request requests.RabRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: "Invalid request data",
			Status:  false,
		})
		return
	}
	if err,validationErr := requests.Validate(request);err != nil {
		resources.BadRequest(ctx, validationErr)
		return
	}
	rab := request.ToModel()
	if err := c.service.Update(rab, uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Message: "RAB updated successfully",
	})
}

func (c *RabController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resources.Response{
			Message: "Invalid ID format",
			Status:  false,
		})
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Message: "RAB deleted successfully",
	})
}
