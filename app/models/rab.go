package models

import (
	"apidanadesa/config"
	"time"

	"gorm.io/gorm"
)
	type Rab struct {
		ID            uint           `gorm:"primaryKey;autoIncrement;"`
		IdKegiatan    uint           `gorm:"not null;"`
		PaketKegiatan   uint        `gorm:"not null" json:"paket_kegiatan"`
		Kode uint         `gorm:"not null;"`
		Anggaran 	float64        `gorm:"type:decimal(20,2);not null;"`
		TahunAnggaran uint   `gorm:"not null;"`
		CreatedAt     time.Time      `gorm:"autoCreateTime"`
		UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
		DeletedAt     gorm.DeletedAt `gorm:"index"`

		// Relations
		// Assuming Kegiatan and TahunAnggaran are other models in your application
		Kegiatan	Kegiatan      `gorm:"foreignKey:IdKegiatan;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		TahunAnggaranData TahunAnggaran `gorm:"foreignKey:TahunAnggaran;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	}
	
	func (Rab) TableName() string {
		return config.GetDBPrefix("rab")
	}
	