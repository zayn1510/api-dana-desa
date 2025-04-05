package models

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type JabatanDesa struct {
	ID        uint           `gorm:"primaryKey"`
	Jabatan   string         `gorm:"type:varchar(100)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (JabatanDesa) TableName() string {
	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal(errenv)
	}
	DB_PREFIX := os.Getenv("DB_PREFIX")
	return DB_PREFIX + "_jabatandesa"
}
