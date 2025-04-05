package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type JabatanDesaService struct {
	db *gorm.DB
}

func NewJabatanDesaService() *JabatanDesaService {
	return &JabatanDesaService{
		db: config.GetDB(),
	}
}
func (s *JabatanDesaService) GetData(offset, limit int) ([]models.JabatanDesa, error) {
	if limit <= 0 {
		limit = 10
	}
	var result []models.JabatanDesa
	err := s.db.Order("id ASC").Offset(offset).Limit(limit).Find(&result).Error
	return result, err
}
func (s *JabatanDesaService) CreateData(r *requests.JabatanDesaRequest) error {
	data := r.ToModelJabatan()
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *JabatanDesaService) UpdateData(r *requests.JabatanDesaRequest, id uint) error {
	data := r.ToModelJabatan()
	var existing models.JabatanDesa
	if err := s.db.First(&existing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID %d tidak ditemukan", id)
		}
		return err
	}
	return s.db.Model(&existing).Updates(data).Error
}
func (s *JabatanDesaService) DeleteData(id uint) error {
	var existing models.JabatanDesa
	if err := s.db.First(&existing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID %d tidak ditemukan", id)
		}
		return err
	}
	return s.db.Delete(&existing).Error
}
