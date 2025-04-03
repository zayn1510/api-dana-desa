package models

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
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
	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal(errenv)
	}
	DB_PREFIX := os.Getenv("DB_PREFIX")
	return DB_PREFIX + "_subbidang"
}
