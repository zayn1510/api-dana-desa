package migrations

	import (
		"gorm.io/gorm"
		"apidanadesa/app/models"
	)

	func MigrateObjekBelanjaDesa(db *gorm.DB) error {
		return db.AutoMigrate(&models.ObjekBelanjaDesa{})
	}
