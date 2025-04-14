package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
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

type Config struct {
	Database struct {
		Hostname string `yaml:"hostname"`
		Root     string `yaml:"root"`
		Password string `yaml:"password"`
		Port     int    `yaml:"port"`
		DBName   string `yaml:"dbname"`
		Prefix   string `yaml:"prefix"`
	} `yaml:"database"`

	JWT struct {
		SecretKey    string `yaml:"secretkey"`
		ExpiredToken int    `yaml:"expiredtoken"`
	} `yaml:"jwt"`
}

var AppConfig Config
var DB_PREFIX string

func init() {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Gagal membaca file config.yaml: %v", err)
	}

	err = yaml.Unmarshal(configFile, &AppConfig)
	if err != nil {
		log.Fatalf("Gagal parsing file config.yaml: %v", err)
	}

	log.Println("config.yaml berhasil diload.")
	DB_PREFIX = AppConfig.Database.Prefix
}
func GetDBPrefix(tablaName string) string {
	return DB_PREFIX + "_" + tablaName
}

func ConnectDB() {
	once.Do(func() {
		if !dbConnected {
			for i := 0; i < 5; i++ {
				dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
					AppConfig.Database.Root,
					AppConfig.Database.Password,
					AppConfig.Database.Hostname,
					AppConfig.Database.Port,
					AppConfig.Database.DBName)

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
