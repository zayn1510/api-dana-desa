package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SumberDanaService struct {
	db *gorm.DB
}

func NewSumberDanaService() *SumberDanaService {
	return &SumberDanaService{
		db: config.GetDB(),
	}
}

func (s *SumberDanaService) IsKodeExist(kode string, id uint) error {
	var count int64
	var model models.SumberDana
	query := s.db.
		Model(&model).
		Where("kode = ?", kode)
	if id > 0 {
		query = query.Where("id != ?", id)
	}
	query.Count(&count)

	if count > 0 {
		return fmt.Errorf("kode Bidang sudah digunakan")
	}
	return nil
}
func (s *SumberDanaService) IsExist(id uint) (models.SumberDana, error) {
	var sumberdana models.SumberDana
	err := s.db.
		Where("id = ?", id).
		First(&sumberdana).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return sumberdana, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return sumberdana, fmt.Errorf("gagal mencari ID")
	}
	return sumberdana, nil
}
func (s *SumberDanaService) GetData(offset, limit int) ([]models.SumberDana, error) {
	var result []models.SumberDana
	err := s.db.Offset(offset).Limit(limit).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *SumberDanaService) CreateData(r *requests.SumberDanaRequestCreate) error {
	duplicate := s.IsKodeExist(r.Kode, 0)
	if duplicate != nil {
		return duplicate
	}
	data := r.ToModel()
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *SumberDanaService) UpdateData(m *requests.SumberDanaRequestUpdate, id uint) error {
	// check model exist by id
	exist, err := s.IsExist(id)
	if err != nil {
		return err
	}

	// check model duplicate by kode,id
	duplicate := s.IsKodeExist(m.Kode, id)
	if duplicate != nil {
		return duplicate
	}
	updates := map[string]interface{}{}
	if m.Kode != "" {
		updates["kode"] = m.Kode
	}
	if m.Keterangan != "" {
		updates["keterangan"] = m.Keterangan
	}
	return s.db.Model(&exist).Updates(updates).Error
}
func (s *SumberDanaService) DeleteData(id uint) error {
	exist, err := s.IsExist(id)
	if err != nil {
		return err
	}
	if err := s.db.Delete(&exist).Error; err != nil {
		return fmt.Errorf("gagal menghapus data: %v", err)
	}
	return nil
}
