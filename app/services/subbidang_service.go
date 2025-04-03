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

type SubBidangService struct {
	db *gorm.DB
}

func NewSubBidangService() *SubBidangService {
	return &SubBidangService{
		db: config.GetDB(),
	}
}

func (s *SubBidangService) CreateSubBidang(req *requests.SubBidangRequestCreate) error {
	data := req.ToModelSubBidang()
	err := s.db.Create(&data).Error
	if err != nil {
		log.Printf("Gagal membuat SubBidang: %v", err)
		return err
	}

	log.Printf("Berhasil membuat SubBidang dengan ID %d", data.ID)
	return nil
}
func (s *SubBidangService) UpdateSubBidang(req *requests.SubBidangRequestCreate, id uint) error {
	var model models.SubBidang
	if err := s.db.First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID %d tidak ditemukan", id)
		}
		return err
	}
	data := req.ToModelSubBidang()
	return s.db.Model(&model).Updates(data).Error
}
func (s *SubBidangService) DeleteSubBidang(id uint) error {
	if err := s.db.First(&models.SubBidang{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID %d tidak ditemukan", id)
		}
		return err
	}
	return s.db.Delete(&models.SubBidang{}, id).Error
}
func (s *SubBidangService) GetAllSubBidang(offset, limit int) ([]models.SubBidang, error) {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	var result []models.SubBidang
	err := s.db.Offset(offset).Limit(limit).Order("id asc").Preload("Bidang", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "kode_bidang", "keterangan")
	}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
