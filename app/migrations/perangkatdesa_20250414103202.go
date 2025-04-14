package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigratePerangkatDesa(db *gorm.DB) error {
		return db.AutoMigrate(&models.PerangkatDesa{})
	}
