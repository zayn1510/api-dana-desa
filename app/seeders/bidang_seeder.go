package main

import (
	"apidanadesa/app/models"
	"apidanadesa/config"
	"fmt"
	"log"
)

func main() {
	db := config.GetDB()
	for i := 101; i <= 10000; i++ {
		bidang := models.Bidang{
			Keterangan: fmt.Sprintf("Bidang Ke-%d", i),
			KodeBidang: fmt.Sprintf("BID%d", i),
		}
		if err := db.Create(&bidang).Error; err != nil {
			log.Fatal(err)
		}
	}
}
