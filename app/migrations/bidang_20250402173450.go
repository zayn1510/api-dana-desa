package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateBidang(db *gorm.DB) error {
		return db.AutoMigrate(&models.Bidang{})
	}
