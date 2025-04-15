package models

import (
	"apidanadesa/config"
	"gorm.io/gorm"
	"time"
)

type PolaKegiatan struct {
	ID         uint           `gorm:"primaryKey;autoIncrement;"`
	Keterangan string         `gorm:"type:varchar(200);"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (PolaKegiatan) TableName() string {
	return config.GetDBPrefix("pola_kegiatan")
}
