package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateJenisPendapatan(db *gorm.DB) error {
		return db.AutoMigrate(&models.JenisPendapatan{})
	}
