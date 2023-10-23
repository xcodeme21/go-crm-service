package migration

import (
	"time"
	"gorm.io/gorm"
)

// Vouchers :nodoc
type Vouchers struct {
    gorm.Model
    VoucherName        string    `gorm:"type:varchar(150); NOT NULL"`
    SeriesId           int       `gorm:"type:integer;unique; NOT NULL"`
    IsExternalVoucher  bool      `gorm:"type:boolean;default:false;"`
    CampaignId         int       `gorm:"type:integer;default:0;"`
    CampaignVoucherId  int       `gorm:"type:integer;default:0;"`
    Point              int64     `gorm:"type:bigint;default:0;"`
    Tier               string    `gorm:"type:varchar(255); NOT NULL"`
    CategoryId         int       `gorm:"type:integer;default:0;"`
    IsLimited          bool      `gorm:"type:boolean;default:false;"`
    StartDate          time.Time `gorm:"type:date;"`
    EndDate            time.Time `gorm:"type:date;"`
    Image              string    `gorm:"type:text; NOT NULL"`
    Description        string    `gorm:"type:text; NOT NULL"`
    Status             bool      `gorm:"type:boolean;default:true;"`
}


// MigrateVouchers :nodoc
func MigrateVouchers(db *gorm.DB) error {
	const tableName = "vouchers"
	// todo to change if need to changing
	const version = "1.1"

	var migrateData Migrate
	var VouchersData Vouchers

	db.Where(&Migrate{Table: tableName}).First(&migrateData)

	// First create installer table
	if migrateData.Table == "" {
		if !db.Migrator().HasTable(&VouchersData) {
			if err := db.AutoMigrate(&Vouchers{}); err != nil {
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
		if err := db.AutoMigrate(&Vouchers{}); err != nil {
			return err // Handle the error, e.g., log it or return it
		}
		migrateData.Version = version
		if err := db.Save(&migrateData).Error; err != nil {
			return err // Handle the error, e.g., log it or return it
		}
	}

	return nil // No errors, return nil
}
