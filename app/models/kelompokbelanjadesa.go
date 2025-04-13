package models

import (
	"apidanadesa/config"
	"gorm.io/gorm"
	"time"
)

type KelompokBelanjaDesa struct {
	ID         uint           `gorm:"primaryKey"`
	Kode       string         `gorm:"type:varchar(10);unique" `
	Keterangan string         `gorm:"type:varchar(100)"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (KelompokBelanjaDesa) TableName() string {
	return config.GetDBPrefix("kelompokbelanjadesa")
}
