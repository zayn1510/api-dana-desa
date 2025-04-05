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

func (s *KegiatanService) IsKodeExist(kode string, id uint) error {
	var count int64
	query := s.db.
		Model(&models.Kegiatan{}).
		Where("kode_kegiatan = ?", kode)
	if id > 0 {
		query = query.Where("id != ?", id)
	}
	query.Count(&count)

	if count > 0 {
		return fmt.Errorf("kode Kegiatan sudah digunakan")
	}
	return nil
}
func (s *KegiatanService) IsExist(id uint) (models.Kegiatan, error) {
	var subBidang models.Kegiatan
	err := s.db.
		Where("id = ?", id).
		First(&subBidang).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return subBidang, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return subBidang, fmt.Errorf("gagal mencari ID")
	}
	return subBidang, nil
}
func (s *KegiatanService) CreateKegiatan(req *requests.KegiatanRequestCreate) error {
	err := s.IsKodeExist(req.KodeKegiatan, 0)
	if err != nil {
		return err
	}
	data := req.ToModelkegiatan()
	err = s.db.Create(&data).Error
	if err != nil {
		log.Printf("Gagal membuat SubBidang: %v", err)
		return err
	}

	log.Printf("Berhasil membuat SubBidang dengan ID %d", data.ID)
	return nil
}
func (s *KegiatanService) UpdateKegiatan(req *requests.KegiatanRequestCreate, id uint) error {

	// check model exist by id
	model, err := s.IsExist(id)
	if err != nil {
		return err
	}
	// check model duplicate by kode,id
	err = s.IsKodeExist(req.KodeKegiatan, id)
	if err != nil {
		return err
	}
	data := req.ToModelkegiatan()
	return s.db.Model(&model).Updates(data).Error
}
func (s *KegiatanService) DeleteKegiatan(id uint) error {
	model, err := s.IsExist(id)
	if err != nil {
		return err
	}
	return s.db.Delete(&model, id).Error
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
