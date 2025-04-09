package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ObjekBelanjaDesaService struct {
	db *gorm.DB
}

func NewObjekBelanjaDesaService() *ObjekBelanjaDesaService {
	return &ObjekBelanjaDesaService{
		db: config.GetDB(),
	}
}

func (s *ObjekBelanjaDesaService) IsKodeExist(kode string, id uint) error {
	var count int64
	query := s.db.Model(&models.ObjekBelanjaDesa{}).Where("kode=?", kode)
	if id > 0 {
		query.Where("id!=?", id)
	}
	query.Count(&count)
	if count > 0 {
		return fmt.Errorf("kode objek belanja sudah digunakan")
	}
	return nil
}

func (s *ObjekBelanjaDesaService) IsExist(id uint) (models.ObjekBelanjaDesa, error) {
	var j models.ObjekBelanjaDesa
	err := s.db.Where("id=?", id).First(&j).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return j, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return j, err
	}
	return j, nil
}

func (s *ObjekBelanjaDesaService) Create(req *requests.ObjekBelanjaDesaRequest) error {
	// check data duplicate by kode,id
	duplicate := s.IsKodeExist(req.Kode, 0)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelObjekBelanja()
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ObjekBelanjaDesaService) Update(id uint, req *requests.ObjekBelanjaDesaRequest) error {
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
	data := req.ToModelObjekBelanja()
	err := s.db.Model(&model).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *ObjekBelanjaDesaService) Delete(id uint) error {
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

func (s *ObjekBelanjaDesaService) GetData(offset, limit int) ([]models.ObjekBelanjaDesa, error) {
	var result []models.ObjekBelanjaDesa
	err := s.db.Offset(offset).Limit(limit).
		Preload("KelompokBelanjaDesa").
		Preload("JenisBelanjaDesa").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
