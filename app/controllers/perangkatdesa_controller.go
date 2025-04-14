package controllers

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}

const maxFileSize = 2 << 20 // 5MB
type PerangkatDesaController struct {
	service *services.PerangkatDesaService
}

func NewPerangkatDesaController() *PerangkatDesaController {
	return &PerangkatDesaController{
		service: services.NewPerangkatDesaService(),
	}
}

func (c *PerangkatDesaController) SubmitForm(ctx *gin.Context) {
	var req requests.PerangkatDesaCreateRequest
	err := ctx.ShouldBindWith(&req, binding.FormMultipart)
	if err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	if err, validationErrors := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validationErrors)
		return
	}
	if req.Foto.Size > maxFileSize {
		resources.BadRequest(ctx, fmt.Errorf("file size too big"))
		return
	}
	ext := strings.ToLower(filepath.Ext(req.Foto.Filename))
	if !allowedExtensions[ext] {
		resources.BadRequest(ctx, fmt.Errorf("file extension not allowed"))
		return
	}
	filePath := fmt.Sprintf("./uploads/%s", req.Foto.Filename)
	layout := "2006-01-02" // Format tanggal: yyyy-mm-dd
	tglLahir, err := time.Parse(layout, req.TglLahir)
	if err != nil {
		resources.BadRequest(ctx, "Format tanggal tidak valid. Gunakan format YYYY-MM-DD")
		return
	}
	tglSk, err := time.Parse(layout, req.TglSK)
	if err != nil {
		resources.BadRequest(ctx, "Format tanggal tidak valid. Gunakan format YYYY-MM-DD")
		return
	}
	data := &models.PerangkatDesa{
		IDJabatan:    req.IDJabatan,
		NamaLengkap:  req.NamaLengkap,
		TempatLahir:  req.TempatLahir,
		TglLahir:     tglLahir,
		TglSK:        tglSk,
		JenisKelamin: req.JenisKelamin,
		NoHandphone:  req.NoHandphone,
		NoSK:         req.NoSK,
		Foto:         filePath,
	}
	if err := c.service.Create(data); err != nil {
		resources.InternalError(ctx, err)
		return
	}
	if err := ctx.SaveUploadedFile(req.Foto, filePath); err != nil {
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "success", req)
}

func (c *PerangkatDesaController) GetAll(ctx *gin.Context) {
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
	bidang := resources.GetPerangkaDesaResource(data)
	ctx.JSON(http.StatusOK, resources.Response{
		Status:  true,
		Message: message,
		Data:    bidang,
	})
}
func (c *PerangkatDesaController) PreviewFile(ctx *gin.Context) {
	filename := ctx.Param("filename")

	// Tentukan path file foto
	filePath := filepath.Join("uploads", filename)

	// Cek apakah file ada
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		resources.NotFound(ctx, fmt.Errorf("file not found"))
		return
	}
	// Kirim file sebagai response
	ctx.File(filePath)
}
