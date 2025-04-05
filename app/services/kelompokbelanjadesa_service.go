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

type KelompokBelanjaDesaService struct {
	db *gorm.DB
}

func NewKelompokBelanjaService() *KelompokBelanjaDesaService {
	return &KelompokBelanjaDesaService{
		db: config.GetDB(),
	}
}
func (s *KelompokBelanjaDesaService) IsKodeExist(kode string, id uint) error {
	var count int64
	query := s.db.
		Model(&models.KelompokBelanjaDesa{}).
		Where("kode= ?", kode)
	if id > 0 {
		query = query.Where("id != ?", id)
	}
	query.Count(&count)

	if count > 0 {
		return fmt.Errorf("kode Kelompok Belanja sudah digunakan")
	}
	return nil
}
func (s *KelompokBelanjaDesaService) IsExist(id uint) (models.KelompokBelanjaDesa, error) {
	var kelompokBelanja models.KelompokBelanjaDesa
	err := s.db.
		Where("id = ?", id).
		First(&kelompokBelanja).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return kelompokBelanja, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return kelompokBelanja, fmt.Errorf("gagal mencari ID")
	}
	return kelompokBelanja, nil
}
func (s *KelompokBelanjaDesaService) CreateData(req *requests.KelompokBelanjaDesaRequest) error {

	duplicate := s.IsKodeExist(req.Kode, 0)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelKelompokBelanja()
	err := s.db.Create(&data).Error
	if err != nil {
		log.Printf("Gagal membuat Kelompok Belanja Desa: %v", err)
		return err
	}

	log.Printf("Berhasil membuat Kelompok Belanja dengan ID %d", data.ID)
	return nil
}
func (s *KelompokBelanjaDesaService) UpdateData(req *requests.KelompokBelanjaDesaRequest, id uint) error {
	model, err := s.IsExist(id)
	if err != nil {
		return err
	}
	duplicate := s.IsKodeExist(req.Kode, id)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelKelompokBelanja()
	return s.db.Model(&model).Updates(data).Error
}
func (s *KelompokBelanjaDesaService) DeleteData(id uint) error {
	model, err := s.IsExist(id)
	if err != nil {
		return err
	}
	return s.db.Delete(&model, id).Error
}
func (s *KelompokBelanjaDesaService) GetAll(offset, limit int) ([]models.KelompokBelanjaDesa, error) {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	var result []models.KelompokBelanjaDesa
	err := s.db.Offset(offset).Limit(limit).Order("id asc").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
