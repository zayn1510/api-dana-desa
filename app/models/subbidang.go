package models

import (
	"apidanadesa/config"
	"gorm.io/gorm"
	"time"
)

type SubBidang struct {
	ID            uint           `gorm:"primaryKey;autoIncrement;"`
	IdBidang      uint           `gorm:"not null;"`
	Keterangan    string         `gorm:"type:varchar(200);"`
	KodeSubBidang string         `gorm:"type:varchar(10);"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	//Relasi
	Bidang Bidang `gorm:"foreignKey:IdBidang" json:"bidang"`
}

func (SubBidang) TableName() string {
	return config.GetDBPrefix("subbidang")
}
