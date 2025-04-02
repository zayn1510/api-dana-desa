package models

import (
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)
	type Bidang struct {
		ID         uint           `gorm:"primaryKey;autoIncrement;"`
		Keterangan string         `gorm:"type:varchar(200);"`
		KodeBidang string         `gorm:"type:varchar(10);"`
		CreatedAt  time.Time      `gorm:"autoCreateTime"`
		UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
		DeletedAt  gorm.DeletedAt `gorm:"index"`
	}
	
	
	func (Bidang) TableName() string {
		errenv := godotenv.Load()
		if errenv != nil {
			log.Fatal(errenv)
		}
		DB_PREFIX := os.Getenv("DB_PREFIX")
		return DB_PREFIX+"_bidang"
	}
	