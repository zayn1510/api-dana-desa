package models

import (
	"apidanadesa/config"
	"time"

	"gorm.io/gorm"
)

type DetailAnggaranKegiatan struct {
	ID uint64 `gorm:"primaryKey" json:"id"`
	IDAnggaranKegiatan uint64  `json:"id_anggaran_kegiatan"`
	NamaPaket string `gorm:"type:varchar(100)" json:"nama_paket"`
	Nilai float64 `json:"nilai"`
	IDPolaKegiatan uint `json:"id_pola_kegiatan"`
	Target string `gorm:"type:varchar(200)"`
	Uraian string `gorm:"type:varchar(200)"`
	Satuan string `gorm:"type:varchar(200)"`
	IDSumberDana uint `json:"id_sumber_dana"`
	SifatKegiatan string `gorm:"type:varchar(200)"`
	LokasiKegiatan string `gorm:"type:varchar(200)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	AnggaranKegiatan AnggaranKegiatan `gorm:"foreignKey:IDAnggaranKegiatan"`
	PolaKegiatan PolaKegiatan `gorm:"foreignKey:IDPolaKegiatan"`
	SumberDana SumberDana `gorm:"foreignKey:IDSumberDana"`
}

func (DetailAnggaranKegiatan) TableName() string {
	return config.GetDBPrefix("detail_anggaran_kegiatan")
}
