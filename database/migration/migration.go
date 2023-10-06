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
func MigrateMigration(db *gorm.DB) {
	var migrateData Migrate
	if db.Migrator().HasTable(&migrateData) == false {
		db.AutoMigrate(&Migrate{})
	}
}

// MigrateExec :nodoc
func MigrateExec(db *gorm.DB) {
	MigrateMigration(db)

	// all table migrate function here
	MigrateVoucherCategories(db)

	log.Println("Auto Migration has been processed")
}
