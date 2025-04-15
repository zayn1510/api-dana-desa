package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigratePolaKegiatan(db *gorm.DB) error {
		return db.AutoMigrate(&models.PolaKegiatan{})
	}
