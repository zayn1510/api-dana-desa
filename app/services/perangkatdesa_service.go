package services

import (
	"apidanadesa/app/models"
	"apidanadesa/config"

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
