package migration

import (
	"gorm.io/gorm"
)

// VoucherCategories :nodoc
type VoucherCategories struct {
	gorm.Model
	Name   string `gorm:"type:varchar(150);unique"`
	Status bool   `gorm:"type:boolean;default:true;"`
}

// MigrateVoucherCategories :nodoc
func MigrateVoucherCategories(db *gorm.DB) error {
	const tableName = "voucher_categories"
	// todo to change if need to changing
	const version = "1.1"

	var migrateData Migrate
	var VoucherCategoriesData VoucherCategories

	db.Where(&Migrate{Table: tableName}).First(&migrateData)

	// First create installer table
	if migrateData.Table == "" {
		if !db.Migrator().HasTable(&VoucherCategoriesData) {
			if err := db.AutoMigrate(&VoucherCategories{}); err != nil {
				return err // Handle the error, e.g., log it or return it
			}
			db.Create(&Migrate{
				Table:   tableName,
				Version: version,
			})
		}
	}

	// Upgrade version
	if migrateData.Version == "1.0" {
		if err := db.AutoMigrate(&VoucherCategories{}); err != nil {
			return err // Handle the error, e.g., log it or return it
		}
		migrateData.Version = version
		if err := db.Save(&migrateData).Error; err != nil {
			return err // Handle the error, e.g., log it or return it
		}
	}

	return nil // No errors, return nil
}
