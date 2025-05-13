package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateSumberDana(db *gorm.DB) error {
		return db.AutoMigrate(&models.SumberDana{})
	}
