package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateObjekPendapatan(db *gorm.DB) error {
		return db.AutoMigrate(&models.ObjekPendapatan{})
	}
