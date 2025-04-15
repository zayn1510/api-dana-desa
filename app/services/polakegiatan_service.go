package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type PolaKegiatanService struct {
	db *gorm.DB
}

func NewPolaKegiatanService() *PolaKegiatanService {
	return &PolaKegiatanService{
		db: config.GetDB(),
	}
}
func (s *PolaKegiatanService) IsExist(id uint) (models.PolaKegiatan, error) {
	var polaKegiatan models.PolaKegiatan
	err := s.db.
		Where("id = ?", id).
		First(&polaKegiatan).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return polaKegiatan, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return polaKegiatan, fmt.Errorf("gagal mencari ID")
	}
	return polaKegiatan, nil
}
func (s *PolaKegiatanService) GetData(offset, limit int) ([]models.PolaKegiatan, error) {
	var result []models.PolaKegiatan
	err := s.db.Offset(offset).Limit(limit).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *PolaKegiatanService) CreateData(r *requests.PolaKegiatanCreate) error {

	data := r.ToModelPolaKegiatan()
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *PolaKegiatanService) UpdateData(r *requests.PolaKegiatanCreate, id uint) error {
	// check model exist by id
	exist, err := s.IsExist(id)
	if err != nil {
		return err
	}
	data := r.ToModelPolaKegiatan()
	return s.db.Model(&exist).Where("id =?", id).Updates(data).Error
}
func (s *PolaKegiatanService) DeleteData(id uint) error {
	exist, err := s.IsExist(id)
	if err != nil {
		return err
	}
	if err := s.db.Delete(&exist).Error; err != nil {
		return fmt.Errorf("gagal menghapus data: %v", err)
	}
	return nil
}
