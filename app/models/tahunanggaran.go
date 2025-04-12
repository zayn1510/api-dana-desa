package models

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
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
	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal(errenv)
	}
	DB_PREFIX := os.Getenv("DB_PREFIX")
	return DB_PREFIX + "_tahunanggaran"
}
