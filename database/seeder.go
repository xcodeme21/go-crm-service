package database

import (
	"fmt"
	"os"
	"time"

	"github.com/xcodeme21/go-crm-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func VoucherCategories() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_TIMEZONE"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	var cek models.VoucherCategories
	db.Table("voucher_categories").First(&cek)
	if cek.ID != 0 {
		fmt.Println("Data found")
	} else {
		// Tambahkan data ke tabel
		for i := 1; i <= 10; i++ {
			name := fmt.Sprintf("Category %d", i)
			db.Create(&models.VoucherCategories{Name: name, ID: i, CreatedAt: time.Now(), UpdatedAt: time.Now()})
		}

		// Menampilkan data
		var products []models.VoucherCategories
		db.Find(&products)
		fmt.Println("Total products:", len(products))
	}
}