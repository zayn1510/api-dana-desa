package controllers

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"mime/multipart"
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
	if err := ctx.ShouldBindJSON(&req); err != nil {
		requests.HandleBindError(ctx, err)
		return
	}

	if err, validationErrors := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validationErrors)
		return
	}
	if err := c.service.CreateData(&req); err != nil {
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

// add function

func (c *BidangController) ImportDataCsv(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	actionStr := ctx.DefaultQuery("action", "0")
	action, err := strconv.Atoi(actionStr)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}
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

	src, err := file.Open()
	if err != nil {
		resources.BadRequest(ctx, "Failed to open uploaded file")
		return
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			resources.InternalError(ctx, err)
		}
	}(src)

	reader := csv.NewReader(src)
	reader.Comma = ','
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		resources.BadRequest(ctx, "Invalid CSV format")
		return
	}
	var preview []resources.BidangCSVRow
	var validModels []models.Bidang
	var failedRows []resources.BidangCSVRow

	for index, record := range records {
		if index == 0 {
			continue
		}
		if len(record) < 2 {
			preview = append(preview, resources.BidangCSVRow{
				Keterangan: "",
				KodeBidang: "",
				Row:        index,
				Error:      "Incomplete row",
			})
		}
		data := requests.BidangRequestCreate{
			Keterangan: strings.TrimSpace(record[1]),
			KodeBidang: strings.TrimSpace(record[0]),
		}

		_, validationErrs := requests.Validate(data)
		var errMsg string
		if validationErrs != nil {
			for _, msg := range validationErrs {
				errMsg = msg
				break
			}
			failedRows = append(failedRows, resources.BidangCSVRow{
				Keterangan: data.Keterangan,
				KodeBidang: data.KodeBidang,
				Row:        index,
				Error:      errMsg,
			})
			continue
		}

		if err := c.service.IsKodeExist(data.KodeBidang, 0); err != nil {
			failedRows = append(failedRows, resources.BidangCSVRow{
				Keterangan: data.Keterangan,
				KodeBidang: data.KodeBidang,
				Row:        index,
				Error:      err.Error(),
			})
			continue
		}
		validModels = append(validModels, data.ToModelCsv())

		preview = append(preview, resources.BidangCSVRow{
			Keterangan: data.Keterangan,
			KodeBidang: data.KodeBidang,
			Row:        index,
			Error:      errMsg,
		})
	}
	if action == 0 {
		offset := (page - 1) * limit

		if offset > len(preview) {
			offset = len(preview)
		}
		end := offset + limit
		if end > len(preview) {
			end = len(preview)
		}
		resources.Success(ctx, "csv preview", preview[offset:end])
		return
	}
	if len(validModels) > 0 {
		err := c.service.BulkInsert(validModels)
		if err != nil {
			resources.InternalError(ctx, err)
			return
		}
	}
	resources.Success(ctx, "berhasil dibuat")
}
