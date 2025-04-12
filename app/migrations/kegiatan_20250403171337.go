package migrations

import (
	"apidanadesa/app/models"
	"gorm.io/gorm"
)

func MigrateKegiatan(db *gorm.DB) error {
	return db.AutoMigrate(&models.Kegiatan{})
}
