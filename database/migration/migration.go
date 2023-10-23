package migration

import (
	"log"

	"gorm.io/gorm"
)

// Migrate :nodoc
type Migrate struct {
	gorm.Model
	Table   string `gorm:"type:varchar(100);unique;not null"`
	Version string `gorm:"type:varchar(10);"`
}

// MigrateMigration :nodoc
func MigrateMigration(db *gorm.DB) error {
	var migrateData Migrate
	if !db.Migrator().HasTable(&migrateData) {
		if err := db.AutoMigrate(&Migrate{}); err != nil {
			return err // Handle the error, e.g., log it or return it
		}
	}
	return nil
}

// MigrateExec :nodoc
func MigrateExec(db *gorm.DB) error {
	if err := MigrateMigration(db); err != nil {
		return err // Propagate the error up if migration fails
	}

	// Call other table migration functions here, e.g., MigrateVoucherCategories(db)
	MigrateVoucherCategories(db)
	MigrateVouchers(db)

	log.Println("Auto Migration has been processed")
	return nil
}
