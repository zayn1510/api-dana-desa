package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type JenisPendapatanService struct {
	db *gorm.DB
}

func NewJenisPendapatanService() *JenisPendapatanService {
	return &JenisPendapatanService{
		db: config.GetDB(),
	}
}

func (s *JenisPendapatanService) IsKodeExist(kode string, id, id_kelompok uint) error {
	var count int64
	query := s.db.Model(&models.JenisPendapatan{}).Where("kode=?", kode).Where("id_kelompok=?", id_kelompok)
	if id > 0 {
		query.Where("id !=?", id)
	}
	query.Count(&count)
	if count > 0 {
		return fmt.Errorf("kode jenis pendapatan sudah digunakan pada id kelompok")
	}
	return nil
}

func (s *JenisPendapatanService) IsExist(id uint) (models.JenisPendapatan, error) {
	var j models.JenisPendapatan
	err := s.db.Where("id=?", id).First(&j).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return j, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return j, err
	}
	return j, nil
}

func (s *JenisPendapatanService) Create(req *requests.JenisPendapatanRequest) error {
	// check data duplicate by kode,id
	duplicate := s.IsKodeExist(req.Kode, 0, req.IdKelompok)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelJenisPendapatan()
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *JenisPendapatanService) Update(id uint, req *requests.JenisPendapatanRequest) error {
	// check data exist by id
	model, exist := s.IsExist(id)
	if exist != nil {
		return exist
	}

	// check data duplicate by kode and id
	duplicate := s.IsKodeExist(req.Kode, id, req.IdKelompok)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelJenisPendapatan()
	err := s.db.Model(&model).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *JenisPendapatanService) Delete(id uint) error {
	// check data exist by id
	model, exist := s.IsExist(id)
	if exist != nil {
		return exist
	}
	err := s.db.Delete(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *JenisPendapatanService) GetData(offset, limit int) ([]models.JenisPendapatan, error) {
	var result []models.JenisPendapatan
	err := s.db.Offset(offset).Limit(limit).Order("id asc").Preload("KelompokPendapatan").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
