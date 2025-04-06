package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type JenisBelanjaDesaService struct {
	db *gorm.DB
}

func NewJenisBelanjaDesaService() *JenisBelanjaDesaService {
	return &JenisBelanjaDesaService{
		db: config.GetDB(),
	}
}

func (s *JenisBelanjaDesaService) IsKodeExist(kode string, id uint) error {
	var count int64
	query := s.db.Model(&models.JenisBelanjaDesa{}).Where("kode=?", kode)
	if id > 0 {
		query.Where("id=?", id)
	}
	query.Count(&count)
	if count > 0 {
		return fmt.Errorf("kode jenis belanja sudah digunakan")
	}
	return nil
}

func (s *JenisBelanjaDesaService) IsExist(id uint) (models.JenisBelanjaDesa, error) {
	var j models.JenisBelanjaDesa
	err := s.db.Where("id=?", id).First(&j).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return j, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return j, err
	}
	return j, nil
}

func (s *JenisBelanjaDesaService) Create(req *requests.JenisBelanjaDesaRequest) error {
	// check data duplicate by kode,id
	duplicate := s.IsKodeExist(req.Kode, 0)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelJenisBelanja()
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *JenisBelanjaDesaService) Update(id uint, req *requests.JenisBelanjaDesaRequest) error {
	// check data exist by id
	model, exist := s.IsExist(id)
	if exist != nil {
		return exist
	}

	// check data duplicate by kode and id
	duplicate := s.IsKodeExist(req.Kode, id)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelJenisBelanja()
	err := s.db.Model(&model).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *JenisBelanjaDesaService) Delete(id uint) error {
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

func (s *JenisBelanjaDesaService) GetData(offset, limit int) ([]models.JenisBelanjaDesa, error) {
	var result []models.JenisBelanjaDesa
	err := s.db.Offset(offset).Limit(limit).Preload("KelompokBelanjaDesa").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
