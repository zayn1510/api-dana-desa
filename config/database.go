package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbconn *gorm.DB
var once sync.Once
var dbConnected bool

var DB_PREFIX string

func init() {
	if _, err := os.Stat(".env"); err == nil {
		log.Println("Loading .env file...")
		if err := godotenv.Load(); err != nil {
			log.Printf("Gagal load .env: %v", err)
		}
	} else {
		log.Println(".env file tidak ditemukan, pakai environment dari container...")
	}
	DB_PREFIX = os.Getenv("DB_PREFIX")
}
func GetDBPrefix(tablaName string) string {
	return DB_PREFIX + "_" + tablaName
}

func ConnectDB() {
	once.Do(func() {
		DB_HOST := os.Getenv("DB_HOST")
		DB_NAME := os.Getenv("DB_NAME")
		DB_USER := os.Getenv("DB_USER")
		DB_PASS := os.Getenv("DB_PASS")
		DB_PORT := os.Getenv("DB_PORT")
		if DB_HOST == "" || DB_NAME == "" || DB_USER == "" || DB_PORT == "" {
			log.Fatal("Pastikan semua variabel database di .env sudah diisi!")
		}

		if !dbConnected {
			for i := 0; i < 5; i++ {
				dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
				database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
					SkipDefaultTransaction: true,
					PrepareStmt:            true,
				})
				if err == nil {
					sqlDB, err := database.DB()
					if err == nil && sqlDB.Ping() == nil {
						dbconn = database
						dbConnected = true
						log.Println("Database berhasil terkoneksi!")
						return
					}
					log.Println("Koneksi berhasil tapi ping gagal, retrying...")
				}
				log.Println("Koneksi Gagal, coba lagi...", err)
				time.Sleep(2 * time.Second)
			}
			log.Fatal("Gagal terkoneksi ke database setelah 5 percobaan")
		} else {
			log.Println("Koneksi database sudah berhasil sebelumnya, tidak mencoba ulang.")
		}
	})
}

func GetDB() *gorm.DB {
	ConnectDB()
	if dbconn == nil {
		log.Fatal("Database is not connected. Pastikan ConnectDB() sudah dipanggil!")
	}
	return dbconn
}
