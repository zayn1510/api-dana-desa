package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateTahunAnggaran(db *gorm.DB) error {
		return db.AutoMigrate(&models.TahunAnggaran{})
	}
