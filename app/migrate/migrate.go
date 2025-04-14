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
	generateSubBidang(db)
	generateKegiatan(db)
	generateTahunAnggaran(db)

	generateJenisBelanjaDesa(db)
	generateKelompokBelanjaDesa(db)
	generateObjekBelanjaDesa(db)
	generateKelompokPendapatan(db)
	generateJenisPendapatan(db)
	generateObjekPendapatan(db)
	generatePerangkatDesa(db)
	generateJabatanDesa(db)
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

func generateSubBidang(db *gorm.DB) {
	fmt.Println("Migrasi tabel SubBidang...")
	if err := migrations.MigrateSubBidang(db); err != nil {
		log.Fatalf("Gagal migrasi SubBidang: %v", err)
	}
	fmt.Println("Migrasi SubBidang selesai.")
}

func generateKegiatan(db *gorm.DB) {
	fmt.Println("Migrasi tabel Kegiatan...")
	if err := migrations.MigrateKegiatan(db); err != nil {
		log.Fatalf("Gagal migrasi Kegiatan: %v", err)
	}
	fmt.Println("Migrasi Kegiatan selesai.")
}

func generateTahunAnggaran(db *gorm.DB) {
	fmt.Println("Migrasi tabel TahunAnggaran...")
	if err := migrations.MigrateTahunAnggaran(db); err != nil {
		log.Fatalf("Gagal migrasi TahunAnggaran: %v", err)
	}
	fmt.Println("Migrasi TahunAnggaran selesai.")
}
func generateJenisBelanjaDesa(db *gorm.DB) {
	fmt.Println("Migrasi tabel JenisBelanjaDesa...")
	if err := migrations.MigrateJenisBelanjaDesa(db); err != nil {
		log.Fatalf("Gagal migrasi JenisBelanjaDesa: %v", err)
	}
	fmt.Println("Migrasi JenisBelanjaDesa selesai.")
}

func generateKelompokBelanjaDesa(db *gorm.DB) {
	fmt.Println("Migrasi tabel KelompokBelanjaDesa...")
	if err := migrations.MigrateKelompokBelanjaDesa(db); err != nil {
		log.Fatalf("Gagal migrasi KelompokBelanjaDesa: %v", err)
	}
	fmt.Println("Migrasi KelompokBelanjaDesa selesai.")
}

func generateObjekBelanjaDesa(db *gorm.DB) {
	fmt.Println("Migrasi tabel ObjekBelanjaDesa...")
	if err := migrations.MigrateObjekBelanjaDesa(db); err != nil {
		log.Fatalf("Gagal migrasi ObjekBelanjaDesa: %v", err)
	}
	fmt.Println("Migrasi ObjekBelanjaDesa selesai.")
}
func generateKelompokPendapatan(db *gorm.DB) {
	fmt.Println("Migrasi tabel KelompokPendapatan...")
	if err := migrations.MigrateKelompokPendapatan(db); err != nil {
		log.Fatalf("Gagal migrasi KelompokPendapatan: %v", err)
	}
	fmt.Println("Migrasi KelompokPendapatan selesai.")
}
func generateJenisPendapatan(db *gorm.DB) {
	fmt.Println("Migrasi tabel JenisPendapatan...")
	if err := migrations.MigrateJenisPendapatan(db); err != nil {
		log.Fatalf("Gagal migrasi JenisPendapatan: %v", err)
	}
	fmt.Println("Migrasi JenisPendapatan selesai.")
}
func generateObjekPendapatan(db *gorm.DB) {
	fmt.Println("Migrasi tabel ObjekPendapatan...")
	if err := migrations.MigrateObjekPendapatan(db); err != nil {
		log.Fatalf("Gagal migrasi ObjekPendapatan: %v", err)
	}
	fmt.Println("Migrasi ObjekPendapatan selesai.")
}
func generatePerangkatDesa(db *gorm.DB) {
	fmt.Println("Migrasi tabel PerangkatDesa...")
	if err := migrations.MigratePerangkatDesa(db); err != nil {
		log.Fatalf("Gagal migrasi PerangkatDesa: %v", err)
	}
	fmt.Println("Migrasi PerangkatDesa selesai.")
}

func generateJabatanDesa(db *gorm.DB) {
	fmt.Println("Migrasi tabel JabatanDesa...")
	if err := migrations.MigrateJabatanDesa(db); err != nil {
		log.Fatalf("Gagal migrasi JabatanDesa: %v", err)
	}
	fmt.Println("Migrasi JabatanDesa selesai.")
}
