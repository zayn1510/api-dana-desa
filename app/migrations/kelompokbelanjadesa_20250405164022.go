package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateKelompokBelanjaDesa(db *gorm.DB) error {
		return db.AutoMigrate(&models.KelompokBelanjaDesa{})
	}
