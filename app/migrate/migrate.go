package main

import (
	"apidanadesa/app/migrations"
	"apidanadesa/config"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func main() {
	config.ConnectDB()
	db := config.GetDB()
	if db == nil {
		log.Fatal("Database tidak terkoneksi!")
	}
	generateBidang(db)
	generateUser(db)
	fmt.Println("Semua migrasi selesai!")
}

func generateBidang(db *gorm.DB) {
	fmt.Println("Migrasi tabel Bidang...")
	if err := migrations.MigrateBidang(db); err != nil {
		log.Fatalf("Gagal migrasi Bidang: %v", err)
	}
	fmt.Println("Migrasi Bidang selesai.")
}
func generateUser(db *gorm.DB) {
	fmt.Println("Migrasi tabel User...")
	if err := migrations.MigrateUser(db); err != nil {
		log.Fatalf("Gagal migrasi User: %v", err)
	}
	fmt.Println("Migrasi User selesai.")
}
