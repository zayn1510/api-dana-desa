package models
	import (
	"github.com/joho/godotenv"
	"os"
	"log"
	)
	type JenisBelanjaDesa struct {
		ID   uint   `gorm:"primaryKey" json:"id"`
		Name string `gorm:"type:varchar(100)" json:"name"`
	}
	
	func (JenisBelanjaDesa) TableName() string {
		errenv := godotenv.Load()
		if errenv != nil {
			log.Fatal(errenv)
		}
		DB_PREFIX := os.Getenv("DB_PREFIX")
		return DB_PREFIX+"_jenisbelanjadesa"
	}
	