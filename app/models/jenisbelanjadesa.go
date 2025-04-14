package models

import (
	"apidanadesa/config"
	"gorm.io/gorm"
	"time"
)

type JenisBelanjaDesa struct {
	ID         uint           `gorm:"primaryKey"`
	IdKelompok uint           `gorm:"not null"`
	Kode       string         `gorm:"type:varchar(10);"`
	Keterangan string         `gorm:"type:varchar(100)"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	//Relasi
	KelompokBelanjaDesa KelompokBelanjaDesa `gorm:"foreignKey:IdKelompok"`
}

func (JenisBelanjaDesa) TableName() string {
	return config.GetDBPrefix("jenis_belanja")
}
