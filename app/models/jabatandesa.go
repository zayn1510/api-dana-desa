package models

import (
	"apidanadesa/config"
	"gorm.io/gorm"
	"time"
)

type JabatanDesa struct {
	ID        uint           `gorm:"primaryKey"`
	Jabatan   string         `gorm:"type:varchar(100)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (JabatanDesa) TableName() string {
	return config.GetDBPrefix("jabatandesa")
}
