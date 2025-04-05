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

type KelompokBelanjaDesaService struct {
	db *gorm.DB
}

func NewKelompokBelanjaService() *KelompokBelanjaDesaService {
	return &KelompokBelanjaDesaService{
		db: config.GetDB(),
	}
}
func (s *KelompokBelanjaDesaService) IsKodeExist(kode string, id uint) (string, bool) {
	var count int64
	s.db.
		Model(&models.KelompokBelanjaDesa{}).
		Where("kode = ? AND id != ? AND deleted_at IS NULL", kode, id).
		Count(&count)

	if count > 0 {
		return fmt.Sprintf("Kode %s sudah digunakan", kode), true
	}
	return "", false
}
func (s *KelompokBelanjaDesaService) CreateData(req *requests.KelompokBelanjaDesaRequest) error {
	data := req.ToModelKelompokBelanja()
	err := s.db.Create(&data).Error
	if err != nil {
		log.Printf("Gagal membuat Kelompok Belanja Desa: %v", err)
		return err
	}

	log.Printf("Berhasil membuat Kelompok Belanja dengan ID %d", data.ID)
	return nil
}
func (s *KelompokBelanjaDesaService) UpdateData(req *requests.KelompokBelanjaDesaRequest, id uint) error {
	var model models.KelompokBelanjaDesa
	if err := s.db.First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID %d tidak ditemukan", id)
		}
		return err
	}
	data := req.ToModelKelompokBelanja()
	return s.db.Model(&model).Updates(data).Error
}
func (s *KelompokBelanjaDesaService) DeleteData(id uint) error {
	if err := s.db.First(&models.KelompokBelanjaDesa{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID %d tidak ditemukan", id)
		}
		return err
	}
	return s.db.Delete(&models.KelompokBelanjaDesa{}, id).Error
}
func (s *KelompokBelanjaDesaService) GetAll(offset, limit int) ([]models.KelompokBelanjaDesa, error) {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	var result []models.KelompokBelanjaDesa
	err := s.db.Offset(offset).Limit(limit).Order("id asc").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
