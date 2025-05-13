package models

import (
	"apidanadesa/config"
	"time"

	"gorm.io/gorm"
)

type SumberDana struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Kode       string         `gorm:"type:varchar(10);not null" json:"kode"`
	Keterangan string         `gorm:"type:varchar(255)" json:"keterangan"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (SumberDana) TableName() string {
	return config.GetDBPrefix("sumber_dana")
}
