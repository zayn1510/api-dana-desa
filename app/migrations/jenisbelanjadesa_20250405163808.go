package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateJenisBelanjaDesa(db *gorm.DB) error {
		return db.AutoMigrate(&models.JenisBelanjaDesa{})
	}
