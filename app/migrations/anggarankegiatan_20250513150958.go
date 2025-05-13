package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateAnggaranKegiatan(db *gorm.DB) error {
		return db.AutoMigrate(&models.AnggaranKegiatan{})
	}
