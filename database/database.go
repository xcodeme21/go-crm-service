package database

import (
	"fmt"
	"log"
	"os"

	"github.com/xcodeme21/go-crm-service/database/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	db, err := Connect()
	if err != nil {
		return nil, err // Return the error if connecting to the database fails
	}

	err = migration.MigrateExec(db)
	if err != nil {
		return nil, err // Return the error if migration fails
	}

	return db, nil
}

// Connect Connection to database
func Connect() (*gorm.DB, error) {
	var (
		dbUser  = os.Getenv("DB_USER")
		dbPass  = os.Getenv("DB_PASSWORD")
		dbHost  = os.Getenv("DB_HOST")
		dbName  = os.Getenv("DB_NAME")
		dbPort  = os.Getenv("DB_PORT")
		TZ      = os.Getenv("DB_TIMEZONE")
		sslMode = os.Getenv("SSL_MODE")
	)

	dsn := fmt.Sprintf(`
		host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s`,
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
		sslMode,
		TZ,
	)
	log.Println("dsn:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Connected to database Failed:", err)
		return nil, err // Return the error if connecting to the database fails
	}
	log.Println("Connected to database")

	return db, nil
}
