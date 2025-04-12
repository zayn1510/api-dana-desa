package models

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type ObjekBelanjaDesa struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	IdKelompok uint           `gorm:"not null;index" json:"id_kelompok"`
	IdJenis    uint           `gorm:"not null;index" json:"id_jenis"`
	Kode       string         `gorm:"type:varchar(10)" json:"kode"`
	Keterangan string         `gorm:"type:varchar(100)" json:"keterangan"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	//Relasi
	KelompokBelanjaDesa KelompokBelanjaDesa `gorm:"foreignKey:IdKelompok"`
	JenisBelanjaDesa    JenisBelanjaDesa    `gorm:"foreignKey:IdJenis"`
}

func (ObjekBelanjaDesa) TableName() string {
	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal(errenv)
	}
	DB_PREFIX := os.Getenv("DB_PREFIX")
	return DB_PREFIX + "_objekbelanjadesa"
}
