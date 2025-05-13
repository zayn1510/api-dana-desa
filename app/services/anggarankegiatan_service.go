package services

import (
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AnggaranKegiatanService struct {
	db *gorm.DB
}

func NewAnggaranKegiatanService() *AnggaranKegiatanService {
	return &AnggaranKegiatanService{
		db: config.GetDB(),
	}
}

func (s *AnggaranKegiatanService) IsExist(id uint) (models.AnggaranKegiatan, error) {
	var anggaran models.AnggaranKegiatan
	err := s.db.
		Where("id = ?", id).
		First(&anggaran).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return anggaran, fmt.Errorf("ID tidak ditemukan")
	} else if err != nil {
		return anggaran, fmt.Errorf("gagal mencari ID")
	}
	return anggaran, nil
}
func (s *AnggaranKegiatanService) GetData(offset, limit int) ([]models.AnggaranKegiatan, error) {
	var result []models.AnggaranKegiatan

	err := s.db.Offset(offset).Limit(limit).
		Preload("Kegiatan", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "keterangan")
		}).
		Preload("PerangkatDesa", func(db *gorm.DB) *gorm.DB {
			return db.
				Select("id", "nama_lengkap", "id_jabatan").
				Preload("JabatanDesa", func(db *gorm.DB) *gorm.DB {
					return db.Select("id", "jabatan")
				})
		}).
		Preload("TahunAnggaran").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *AnggaranKegiatanService) CreateData(r *requests.AnggaranKegiatanRequestCreate) error {
	data := r.ToModel()
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *AnggaranKegiatanService) UpdateData(r *requests.AnggaranKegiatanRequestUpdate, id uint) error {
	// check model exist by id
	exist, err := s.IsExist(id)
	if err != nil {
		return err
	}

	data := r.ToModelUpdate()
	return s.db.Model(&exist).Updates(data).Error
}

func (s *AnggaranKegiatanService) DeleteData(id uint) error {
	exist, err := s.IsExist(id)
	if err != nil {
		return err
	}
	if err := s.db.Delete(&exist).Error; err != nil {
		return fmt.Errorf("gagal menghapus data: %v", err)
	}
	return nil
}
