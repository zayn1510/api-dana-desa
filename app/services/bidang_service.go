package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BidangService struct {
	db *gorm.DB
}

func NewBidangService() *BidangService {
	return &BidangService{
		db: config.GetDB(),
	}
}

func (s *BidangService) IsKodeExist(kode string, id uint) error {
	var count int64
	query := s.db.
		Model(&models.Bidang{}).
		Where("kode_bidang = ?", kode)
	if id > 0 {
		query = query.Where("id != ?", id)
	}
	query.Count(&count)

	if count > 0 {
		return fmt.Errorf("kode Bidang sudah digunakan")
	}
	return nil
}
func (s *BidangService) IsExist(id uint) (models.Bidang, error) {
	var bidang models.Bidang
	err := s.db.
		Where("id = ?", id).
		First(&bidang).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return bidang, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return bidang, fmt.Errorf("gagal mencari ID")
	}
	return bidang, nil
}
func (s *BidangService) GetData(offset, limit int) ([]models.Bidang, error) {
	var result []models.Bidang
	err := s.db.Offset(offset).Limit(limit).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *BidangService) CreateData(r *requests.BidangRequestCreate) error {
	duplicate := s.IsKodeExist(r.KodeBidang, 0)
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
func (s *BidangService) UpdateData(r *requests.BidangRequestCreate, id uint) error {
	// check model exist by id
	exist, err := s.IsExist(id)
	if err != nil {
		return err
	}

	// check model duplicate by kode,id
	duplicate := s.IsKodeExist(r.KodeBidang, id)
	if duplicate != nil {
		return duplicate
	}
	data := r.ToModel()
	return s.db.Model(&exist).Where("id =?", id).Updates(data).Error
}
func (s *BidangService) DeleteData(id uint) error {
	exist, err := s.IsExist(id)
	if err != nil {
		return err
	}
	if err := s.db.Delete(&exist).Error; err != nil {
		return fmt.Errorf("Gagal menghapus data: %v", err)
	}
	return nil
}
