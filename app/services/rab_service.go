package services

import (
	"apidanadesa/app/models"
	"apidanadesa/config"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type RabService struct {
	db *gorm.DB
}

func NewRabService() *RabService {
	return &RabService{
		db: config.GetDB(),
	}
}

func (s *RabService) Create(rab *models.Rab) error {
	return s.db.Create(rab).Error
}

func (s *RabService) IsExist(id uint) (models.Rab, error) {
	var rab models.Rab
	err := s.db.
		Where("id = ?", id).
		First(&rab).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rab, fmt.Errorf("ID tidak ditemukan")
		}
		return rab, fmt.Errorf("gagal mencari ID: %v", err)
	}
	return rab, nil
}
func (s *RabService) GetData(offset, limit int) ([]models.Rab, error) {
	var result []models.Rab
	err := s.db.Offset(offset).Limit(limit).Find(&result).Error
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data: %v", err)
	}
	return result, nil
}
func (s *RabService) Update(rab *models.Rab, id uint) error {
	existingRab, err := s.IsExist(id)
	if err != nil {
		return err
	}
	if existingRab.ID == 0 {
		return fmt.Errorf("data RAB dengan ID %d tidak ditemukan", id)
	}
	return s.db.Model(&existingRab).Updates(rab).Error
}

func (s *RabService) Delete(id uint) error {
	existingRab, err := s.IsExist(id)
	if err != nil {
		return err
	}

	if existingRab.ID == 0 {
		return fmt.Errorf("data RAB dengan ID %d tidak ditemukan", id)
	}
	return s.db.Delete(&models.Rab{}, id).Error
}
