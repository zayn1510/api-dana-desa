package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ObjekPendapatanService struct {
	db *gorm.DB
}

func NewObjekPendapatanService() *ObjekPendapatanService {
	return &ObjekPendapatanService{
		db: config.GetDB(),
	}
}

func (s *ObjekPendapatanService) IsKodeExist(kode string, id, id_kelompok, id_jenis uint) error {
	var count int64
	query := s.db.Model(&models.ObjekPendapatan{}).
		Where("kode=?", kode).
		Where("id_kelompok=?", id_kelompok).
		Where("id_jenis=?", id_jenis)
	if id > 0 {
		query.Where("id!=?", id)
	}
	query.Count(&count)
	if count > 0 {
		return fmt.Errorf("kode objek pendapatan sudah digunakan pada id kelompok dan jenis yang sama")
	}
	return nil
}

func (s *ObjekPendapatanService) IsExist(id uint) (models.ObjekPendapatan, error) {
	var j models.ObjekPendapatan
	err := s.db.Where("id=?", id).First(&j).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return j, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return j, err
	}
	return j, nil
}

func (s *ObjekPendapatanService) Create(req *requests.ObjekPendapatanRequest) error {
	// check data duplicate by kode,id
	duplicate := s.IsKodeExist(req.Kode, 0, req.IdKelompok, req.IdJenis)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelObjekPendapatan()
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ObjekPendapatanService) Update(id uint, req *requests.ObjekPendapatanRequest) error {
	// check data exist by id
	model, exist := s.IsExist(id)
	if exist != nil {
		return exist
	}

	// check data duplicate by kode and id
	duplicate := s.IsKodeExist(req.Kode, id, req.IdKelompok, req.IdJenis)
	if duplicate != nil {
		return duplicate
	}
	data := req.ToModelObjekPendapatan()
	err := s.db.Model(&model).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *ObjekPendapatanService) Delete(id uint) error {
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

func (s *ObjekPendapatanService) GetData(offset, limit int) ([]models.ObjekPendapatan, error) {
	var result []models.ObjekPendapatan
	err := s.db.Offset(offset).Limit(limit).
		Preload("KelompokPendapatan").
		Preload("JenisPendapatan").
		Order("id asc").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
