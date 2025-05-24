package services

import (
	"apidanadesa/app/models"
	"apidanadesa/config"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type DetailAnggaranKegiatanService struct {
	db *gorm.DB
}

func NewDetailAnggaranKegiatanService() *DetailAnggaranKegiatanService {
	return &DetailAnggaranKegiatanService{
		db: config.GetDB(),
	}
}

func (s *DetailAnggaranKegiatanService) IsExist(id uint64) (models.DetailAnggaranKegiatan, error) {
	var exist models.DetailAnggaranKegiatan
	err := s.db.First(&exist, id).Error
	if err != nil {
		return exist, err
	}
	return exist, nil
}

// Crud function
func (s *DetailAnggaranKegiatanService) GetAll(offset, limit int) ([]models.DetailAnggaranKegiatan, error) {

	var results []models.DetailAnggaranKegiatan
	err := s.db.
		Limit(limit).
		Offset(offset).
		Order("id ASC").
		Find(&results).Error

	return results, err
}

func (s *DetailAnggaranKegiatanService) Save(data *models.DetailAnggaranKegiatan) error {
	return s.db.Create(data).Error
}

func (s *DetailAnggaranKegiatanService) Update(data *models.DetailAnggaranKegiatan, id uint64) error {
	exist, err := s.IsExist(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID tidak ditemukan")
		}
		return err
	}

	return s.db.Model(&exist).Updates(data).Error
}

func (s *DetailAnggaranKegiatanService) Delete(id uint64) error {
	exist, err := s.IsExist(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID tidak ditemukan")
		}
		return err
	}
	return s.db.Delete(&exist).Error
}
