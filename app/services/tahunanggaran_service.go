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

func (s *TahunAnggaranService) IsTahunExist(tahun string, id uint) error {
	var count int64
	query := s.db.
		Model(&models.TahunAnggaran{}).
		Where("tahun = ?", tahun)
	if id > 0 {
		query = query.Where("id != ?", id)
	}
	query.Count(&count)

	if count > 0 {
		return fmt.Errorf("tahun ini sudah digunakan")
	}
	return nil
}
func (s *TahunAnggaranService) IsExist(id uint) (models.TahunAnggaran, error) {
	var tahun models.TahunAnggaran
	err := s.db.
		Where("id = ?", id).
		First(&tahun).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tahun, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return tahun, fmt.Errorf("gagal mencari ID")
	}
	return tahun, nil
}
func (s *TahunAnggaranService) CreateTahunAnggaran(req *requests.TahunAnggaranRequest) error {
	err := s.IsTahunExist(req.Tahun, 0)
	if err != nil {
		return err
	}
	data := req.ToModelTahunAnggaran()
	err = s.db.Create(&data).Error
	if err != nil {
		log.Printf("Gagal membuat Kegiatan: %v", err)
		return err
	}

	log.Printf("Berhasil membuat Kegiatan dengan ID %d", data.ID)
	return nil
}
func (s *TahunAnggaranService) UpdateTahunAnggaran(req *requests.TahunAnggaranRequest, id uint) error {

	// check model exist by id
	model, err := s.IsExist(id)
	if err != nil {
		return err
	}

	// check model duplicate by tahun,id
	err = s.IsTahunExist(req.Tahun, id)
	if err != nil {
		return err
	}
	updateData := map[string]interface{}{
		"tahun":  req.Tahun,
		"status": req.Status,
	}
	return s.db.Model(&model).Updates(updateData).Error
}
func (s *TahunAnggaranService) DeleteTahunAnggaran(id uint) error {
	model, err := s.IsExist(id)
	if err != nil {
		return err
	}
	return s.db.Delete(&model, id).Error
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
