package models

import (
	"apidanadesa/config"
	"time"

	"gorm.io/gorm"
)

type AnggaranKegiatan struct {
	ID              uint64         `gorm:"primaryKey" json:"id"`
	IDKegiatan      int            `json:"id_kegiatan"`
	Lokasi          string         `gorm:"type:varchar(100)" json:"lokasi"`
	Waktu           string         `gorm:"type:varchar(20)" json:"waktu"`
	IDPerangkatDesa int            `json:"id_perangkat_desa"`
	Keluaran        string         `gorm:"type:varchar(200)" json:"keluaran"`
	Volume          string         `gorm:"type:varchar(50)" json:"volume"`
	Pagu            float64        `json:"pagu"`
	IDTahunAnggaran int            `json:"id_tahun_anggaran"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	// Relation
	Kegiatan      Kegiatan      `gorm:"foreignKey:IDKegiatan"`
	PerangkatDesa PerangkatDesa `gorm:"foreignKey:IDPerangkatDesa"`
	TahunAnggaran TahunAnggaran `gorm:"foreignKey:IDTahunAnggaran"`
}

func (AnggaranKegiatan) TableName() string {
	return config.GetDBPrefix("anggaran_kegiatan")
}
