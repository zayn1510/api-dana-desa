package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigratedetailAnggaranKegiatan(db *gorm.DB) error {
		return db.AutoMigrate(&models.DetailAnggaranKegiatan{})
	}
