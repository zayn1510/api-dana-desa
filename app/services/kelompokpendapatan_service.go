package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type KelompokPendapatanService struct {
	db *gorm.DB
}

func NewKelompokPendapatanService() *KelompokPendapatanService {
	return &KelompokPendapatanService{
		db: config.GetDB(),
	}
}

func (s *KelompokPendapatanService) IsKodeExist(kode string, id uint) error {
	var count int64
	query := s.db.Model(&models.KelompokPendapatan{}).Where("kode =?", kode)
	if id > 0 {
		query = query.Where("id !=?", id)
	}
	query.Count(&count)
	if count > 0 {
		return fmt.Errorf("kode pendapatan : %s sudah digunakan", kode)
	}
	return nil
}

func (s *KelompokPendapatanService) IsExist(id uint) (models.KelompokPendapatan, error) {
	var kelompokPendapatan models.KelompokPendapatan
	err := s.db.
		Where("id = ?", id).
		First(&kelompokPendapatan).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return kelompokPendapatan, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return kelompokPendapatan, fmt.Errorf("gagal mencari ID")
	}
	return kelompokPendapatan, nil
}
func (s *KelompokPendapatanService) GetAll(offset, limit int) ([]models.KelompokPendapatan, error) {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	var results []models.KelompokPendapatan
	err := s.db.Offset(offset).Limit(limit).Order("id asc").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (s *KelompokPendapatanService) Created(req *requests.KelompokPendapatanBelanja) error {
	// check kode duplicate
	duplicate := s.IsKodeExist(req.Kode, 0)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelKelompokPendapatan()
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *KelompokPendapatanService) UpdateData(req *requests.KelompokPendapatanBelanja, id uint) error {
	model, err := s.IsExist(id)
	if err != nil {
		return err
	}
	duplicate := s.IsKodeExist(req.Kode, id)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelKelompokPendapatan()
	return s.db.Model(&model).Updates(data).Error
	return nil
}
func (s *KelompokPendapatanService) DeleteData(id uint) error {
	model, err := s.IsExist(id)
	if err != nil {
		return err
	}
	return s.db.Delete(&model, id).Error
}

func (s *KelompokPendapatanService) BulkInsert(data []models.KelompokPendapatan) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return tx.Debug().CreateInBatches(&data, 1000).Error
	})
}
