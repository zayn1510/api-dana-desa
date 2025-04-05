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

func (s *SubBidangService) IsKodeExist(kode string, id uint) error {
	var count int64
	query := s.db.
		Model(&models.SubBidang{}).
		Where("kode_sub_bidang = ?", kode)
	if id > 0 {
		query = query.Where("id != ?", id)
	}
	query.Count(&count)

	if count > 0 {
		return fmt.Errorf("Kode Sub Bidang sudah digunakan")
	}
	return nil
}
func (s *SubBidangService) IsExist(id uint) (models.SubBidang, bool, error) {
	var subBidang models.SubBidang
	err := s.db.
		Where("id = ?", id).
		First(&subBidang).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return subBidang, false, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return subBidang, false, fmt.Errorf("Gagal mencari ID")
	}
	return subBidang, true, nil
}

func (s *SubBidangService) CreateSubBidang(req *requests.SubBidangRequestCreate) error {
	// check kode bidang duplicate for update
	err := s.IsKodeExist(req.KodeBidang, 0)
	if err != nil {
		return err
	}
	data := req.ToModelSubBidang()
	err = s.db.Create(&data).Error
	if err != nil {
		log.Printf("Gagal membuat SubBidang: %v", err)
		return err
	}

	log.Printf("Berhasil membuat SubBidang dengan ID %d", data.ID)
	return nil
}
func (s *SubBidangService) UpdateSubBidang(req *requests.SubBidangRequestCreate, id uint) error {

	// check id exist
	model, exist, err := s.IsExist(id)
	if err != nil {
		return err
	}
	if !exist {
		return err
	}

	// check kode bidang duplicate for update
	err = s.IsKodeExist(req.KodeBidang, id)
	if err != nil {
		return err
	}
	data := req.ToModelSubBidang()
	return s.db.Model(&model).Updates(data).Error

}
func (s *SubBidangService) DeleteSubBidang(id uint) error {
	model, exist, err := s.IsExist(id)
	if err != nil {
		return err
	}
	if !exist {
		return err
	}
	return s.db.Delete(&model, id).Error
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
