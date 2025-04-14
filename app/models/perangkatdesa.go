package models

import (
	"apidanadesa/config"
	"gorm.io/gorm"
	"time"
)

type PerangkatDesa struct {
	ID           uint      `gorm:"primaryKey"`
	IDJabatan    uint      `gorm:"not null"`
	NamaLengkap  string    `gorm:"type:varchar(100)"`
	TempatLahir  string    `gorm:"type:varchar(100)"`
	TglLahir     time.Time `gorm:"type:date"`
	JenisKelamin string    `gorm:"type:enum('Laki-laki','Perempuan')"`
	NoSK         string    `gorm:"type:varchar(100);not null"`
	TglSK        time.Time `gorm:"type:date;not null"`
	NoHandphone  string    `gorm:"type:varchar(20);not null"`
	Foto         string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt

	//Relasi
	JabatanDesa JabatanDesa `gorm:"foreignKey:IDJabatan"`
}

func (PerangkatDesa) TableName() string {
	return config.GetDBPrefix("perangkat_desa")
}
