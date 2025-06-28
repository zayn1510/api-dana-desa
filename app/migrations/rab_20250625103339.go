package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func Migraterab(db *gorm.DB) error {
		return db.AutoMigrate(&models.Rab{})
	}
