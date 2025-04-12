package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateJabatanDesa(db *gorm.DB) error {
		return db.AutoMigrate(&models.JabatanDesa{})
	}
