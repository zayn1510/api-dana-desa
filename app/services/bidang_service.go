package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"

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

func (s *BidangService) GetData(offset, limit int) ([]models.Bidang, error) {
	var result []models.Bidang
	err := s.db.Offset(offset).Limit(limit).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *BidangService) CreateData(r *requests.BidangRequestCreate) error {
	data := r.ToModel()
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *BidangService) UpdateData(r *requests.BidangRequestCreate, id uint) error {
	data := r.ToModel()
	err := s.db.Model(&data).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *BidangService) DeleteData(id uint) error {
	err := s.db.Where("id = ?", id).Delete(&models.Bidang{}).Error
	if err != nil {
		return err
	}
	return nil
}
