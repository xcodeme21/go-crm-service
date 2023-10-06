package migration

import (
	"gorm.io/gorm"
)

// VoucherCategories :nodoc
type VoucherCategories struct {
	gorm.Model
	Name               string `gorm:"type:varchar(150);unique"`
	// Status                  bool   `gorm:"type:boolean;default:true;"`
}

// MigrateVoucherCategories :nodoc
func MigrateVoucherCategories(db *gorm.DB) {
	const tableName = "voucher_categories"
	// todo to change if need to changing
	const version = "1.1"

	var migrateData Migrate
	var VoucherCategoriesData VoucherCategories

	db.Where(&Migrate{Table: tableName}).First(&migrateData)

	// First create installer table
	if migrateData.Table == "" {
		if !db.Migrator().HasTable(&VoucherCategoriesData) {
			db.AutoMigrate(&VoucherCategories{})
			db.Create(&Migrate{
				Table:   tableName,
				Version: version,
			})
		}
	}

	// Upgrade version
	if migrateData.Version == "1.0" {
		db.AutoMigrate(&VoucherCategories{})
		migrateData.Version = version
		db.Save(&migrateData)
	}

}