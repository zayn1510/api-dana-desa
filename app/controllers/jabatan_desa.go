package controllers

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
	"strings"
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
		resources.BadRequest(ctx, err)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	offset := (page - 1) * limit
	data, err := c.service.GetData(offset, limit)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	message := "Data kosong"
	if len(data) > 0 {
		message = "data berhasil dimuat"
	}
	response := resources.GetResponseJabatanDesa(data)
	resources.Success(ctx, message, response)

}

func (c *JabatanDesaController) CreateData(ctx *gin.Context) {
	var req requests.JabatanDesaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		requests.HandleBindError(ctx, err)
		return
	}
	if err, validation := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validation)
		return
	}

	err := c.service.CreateData(&req)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "data berhasil dimuat", req)
}
func (c *JabatanDesaController) UpdateData(ctx *gin.Context) {

	var req requests.JabatanDesaRequest
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

	err = c.service.UpdateData(&req, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			resources.NotFound(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "data berhasil diupdate", req)
}
func (c *JabatanDesaController) DeleteData(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	err = c.service.DeleteData(uint(id))
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

func (c *JabatanDesaController) ImportDataCsv(ctx *gin.Context) {
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
	var preview []resources.JabatanCSVRow
	var validModels []models.JabatanDesa
	var failedRows []resources.JabatanCSVRow

	for index, record := range records {
		if index == 0 {
			continue
		}
		if len(record) < 1 {
			preview = append(preview, resources.JabatanCSVRow{
				Jabatan: "",
				Row:     index,
				Error:   "Incomplete row",
			})
		}
		data := requests.JabatanDesaRequest{
			Jabatan: strings.TrimSpace(record[0]),
		}

		_, validationErrs := requests.Validate(data)
		var errMsg string
		if validationErrs != nil {
			for _, msg := range validationErrs {
				errMsg = msg
				break
			}
			failedRows = append(failedRows, resources.JabatanCSVRow{
				Jabatan: data.Jabatan,
				Row:     index,
				Error:   errMsg,
			})
			continue
		}
		validModels = append(validModels, data.ToModelJabatanCsv())

		preview = append(preview, resources.JabatanCSVRow{
			Jabatan: data.Jabatan,
			Row:     index,
			Error:   errMsg,
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
