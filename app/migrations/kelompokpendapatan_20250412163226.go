package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateKelompokPendapatan(db *gorm.DB) error {
		return db.AutoMigrate(&models.KelompokPendapatan{})
	}
