package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateSubBidang(db *gorm.DB) error {
		return db.AutoMigrate(&models.SubBidang{})
	}
