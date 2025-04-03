package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type TahunAnggaranService struct {
	db *gorm.DB
}

func NewTahunAnggaranService() *TahunAnggaranService {
	return &TahunAnggaranService{
		db: config.GetDB(),
	}
}

func (s *TahunAnggaranService) CreateTahunAnggaran(req *requests.TahunAnggaranRequest) error {
	data := req.ToModelTahunAnggaran()
	err := s.db.Create(&data).Error
	if err != nil {
		log.Printf("Gagal membuat Kegiatan: %v", err)
		return err
	}

	log.Printf("Berhasil membuat Kegiatan dengan ID %d", data.ID)
	return nil
}
func (s *TahunAnggaranService) UpdateTahunAnggaran(req *requests.TahunAnggaranRequest, id uint) error {
	var model models.TahunAnggaran
	if err := s.db.First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID %d tidak ditemukan", id)
		}
		return err
	}
	updateData := map[string]interface{}{
		"tahun":  req.Tahun,
		"status": req.Status,
	}
	return s.db.Model(&model).Updates(updateData).Error
}
func (s *TahunAnggaranService) DeleteTahunAnggaran(id uint) error {
	if err := s.db.First(&models.TahunAnggaran{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID %d tidak ditemukan", id)
		}
		return err
	}
	return s.db.Delete(&models.TahunAnggaran{}, id).Error
}
func (s *TahunAnggaranService) GetAllTahun(offset, limit int) ([]models.TahunAnggaran, error) {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	var result []models.TahunAnggaran
	err := s.db.Offset(offset).Limit(limit).Order("id asc").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
