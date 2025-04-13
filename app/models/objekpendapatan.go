package models

import (
	"apidanadesa/config"
	"gorm.io/gorm"
	"time"
)

type ObjekPendapatan struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	IdKelompok uint           `gorm:"not null;index" json:"id_kelompok"`
	IdJenis    uint           `gorm:"not null;index" json:"id_jenis"`
	Kode       string         `gorm:"type:varchar(10)" json:"kode"`
	Keterangan string         `gorm:"type:varchar(100)" json:"keterangan"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	//Relasi
	KelompokPendapatan KelompokPendapatan `gorm:"foreignKey:IdKelompok"`
	JenisPendapatan    JenisPendapatan    `gorm:"foreignKey:IdJenis"`
}

func (ObjekPendapatan) TableName() string {
	return config.GetDBPrefix("objekpendapatan")
}
