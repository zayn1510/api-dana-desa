package models

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type JenisBelanjaDesa struct {
	ID         uint           `gorm:"primaryKey"`
	IdKelompok uint           `gorm:"not null"`
	Kode       string         `gorm:"type:varchar(10);"`
	Keterangan string         `gorm:"type:varchar(100)"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	//Relasi
	KelompokBelanjaDesa KelompokBelanjaDesa `gorm:"foreignKey:IdKelompok"`
}

func (JenisBelanjaDesa) TableName() string {
	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal(errenv)
	}
	DB_PREFIX := os.Getenv("DB_PREFIX")
	return DB_PREFIX + "_jenisbelanjadesa"
}
