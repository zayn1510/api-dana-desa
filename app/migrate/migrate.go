package main

import (
	"apidanadesa/app/migrations"
	"apidanadesa/config"
	"fmt"
	"log"
	"gorm.io/gorm"
)

func main() {
	config.ConnectDB()
	db := config.GetDB()
	if db == nil {
		log.Fatal("Database tidak terkoneksi!")
	}
		generateBidang(db)
	fmt.Println("Semua migrasi selesai!")
}

func generateBidang(db *gorm.DB) {
		fmt.Println("Migrasi tabel Bidang...")
		if err := migrations.MigrateBidang(db); err != nil {
			log.Fatalf("Gagal migrasi Bidang: %v", err)
		}
		fmt.Println("Migrasi Bidang selesai.")
}