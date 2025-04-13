package models

import (
	"apidanadesa/config"
	"gorm.io/gorm"
	"time"
)

type Bidang struct {
	ID         uint           `gorm:"primaryKey;autoIncrement;"`
	Keterangan string         `gorm:"type:varchar(200);"`
	KodeBidang string         `gorm:"type:varchar(10);"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (Bidang) TableName() string {
	return config.GetDBPrefix("bidang")
}
