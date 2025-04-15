package services

import (
	"apidanadesa/app/models"
	"apidanadesa/config"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type PerangkatDesaService struct {
	db *gorm.DB
}

func NewPerangkatDesaService() *PerangkatDesaService {
	return &PerangkatDesaService{
		db: config.GetDB(),
	}
}

func (s *PerangkatDesaService) GetData(offset, limit int) ([]models.PerangkatDesa, error) {
	var result []models.PerangkatDesa
	err := s.db.Offset(offset).Limit(limit).
		Preload("JabatanDesa").
		Order("id asc").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PerangkatDesaService) Create(data *models.PerangkatDesa) error {
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PerangkatDesaService) IsExist(id uint) (models.PerangkatDesa, error) {
	var perangkatDesa models.PerangkatDesa
	err := s.db.Where("id = ?", id).First(&perangkatDesa).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return perangkatDesa, fmt.Errorf("ID not found")
	} else if err != nil {
		return perangkatDesa, err
	}
	return perangkatDesa, nil
}

func (s *PerangkatDesaService) Update(data *models.PerangkatDesa, id uint) error {
	perangkatDesa, err := s.IsExist(id)
	if err != nil {
		return err
	}
	if data.Foto == "" {
		data.Foto = perangkatDesa.Foto
	}
	err = s.db.Model(&perangkatDesa).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PerangkatDesaService) Delete(id uint) (models.PerangkatDesa, error) {
	perangkatDesa, err := s.IsExist(id)
	if err != nil {
		return perangkatDesa, err
	}
	err = s.db.Delete(&perangkatDesa).Error
	if err != nil {
		return perangkatDesa, err
	}
	return perangkatDesa, nil
}
