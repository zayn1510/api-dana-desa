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

type KegiatanService struct {
	db *gorm.DB
}

func NewKegiatanService() *KegiatanService {
	return &KegiatanService{
		db: config.GetDB(),
	}
}
func (s *KegiatanService) CreateKegiatan(req *requests.KegiatanRequestCreate) error {
	data := req.ToModelkegiatan()
	err := s.db.Create(&data).Error
	if err != nil {
		log.Printf("Gagal membuat SubBidang: %v", err)
		return err
	}

	log.Printf("Berhasil membuat SubBidang dengan ID %d", data.ID)
	return nil
}
func (s *KegiatanService) UpdateKegiatan(req *requests.KegiatanRequestCreate, id uint) error {
	var model models.Kegiatan
	if err := s.db.First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID %d tidak ditemukan", id)
		}
		return err
	}
	data := req.ToModelkegiatan()
	return s.db.Model(&model).Updates(data).Error
}
func (s *KegiatanService) DeleteKegiatan(id uint) error {
	if err := s.db.First(&models.Kegiatan{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID %d tidak ditemukan", id)
		}
		return err
	}
	return s.db.Delete(&models.Kegiatan{}, id).Error
}
func (s *KegiatanService) GetAllKegiatan(offset, limit int) ([]models.Kegiatan, error) {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	var result []models.Kegiatan
	err := s.db.Offset(offset).Limit(limit).Order("id asc").Preload("Bidang", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "kode_bidang", "keterangan")
	}).Preload("SubBidang", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "kode_sub_bidang", "keterangan")
	}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
