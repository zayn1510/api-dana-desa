package models

import (
	"apidanadesa/config"
	"gorm.io/gorm"
	"time"
)

type TahunAnggaran struct {
	ID        uint           `gorm:"primary_key;auto_increment"`
	Tahun     string         `gorm:"type:varchar(30);unique_index"`
	Status    int            `gorm:"type:tinyint(1);default:0;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (TahunAnggaran) TableName() string {
	return config.GetDBPrefix("tahunanggaran")
}
