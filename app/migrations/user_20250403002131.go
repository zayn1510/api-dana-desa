package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateUser(db *gorm.DB) error {
		return db.AutoMigrate(&models.User{})
	}
