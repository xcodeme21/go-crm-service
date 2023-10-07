package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xcodeme21/go-crm-service/api/cms/controllers"
	"github.com/xcodeme21/go-crm-service/api/cms/providers"
	"github.com/xcodeme21/go-crm-service/api/cms/services"
	"github.com/xcodeme21/go-crm-service/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func lostInSpce(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":        c.Writer.Status(),
		"data":          nil,
		"error_message": "Lost in space",
	})
}

func createDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_TIMEZONE"))
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DB_NAME"))
	DB.Exec(createDatabaseCommand)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   false,
			"message": "Yayyyy I'am Gin Gonic",
		})
	})

	// cors configuration
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"OPTIONS", "PUT", "POST", "GET", "DELETE", "PATCH"}
	r.Use(cors.New(config))

	createDatabase()

	// Initializes databaseSource
	db, _ := database.Initialize()
	r.Use(database.Inject(db))

	//Seeder
	database.VoucherCategories()

	//Connection
	cn, _ := database.Connect()

	voucherCategoriesService := services.NewVoucherCategoriesService(providers.NewDBVoucherCategoriesProvider(cn))

	// Dereference the pointer to pass the underlying value to the constructor
	categoriesController := controllers.NewVoucherCategoriesController(*voucherCategoriesService)

	cmsGroup := r.Group("/api/cms/")
	cmsGroup.GET("/voucher-categories", categoriesController.ListCategories)
	cmsGroup.GET("/voucher-categories/:id", categoriesController.DetailCategory)
	cmsGroup.POST("/voucher-categories", categoriesController.CreateCategory)
	cmsGroup.PUT("/voucher-categories/:id", categoriesController.UpdateCategory)
	cmsGroup.DELETE("/voucher-categories/:id", categoriesController.DeleteCategory)

	port := os.Getenv("PORT")
	r.NoRoute(lostInSpce)
	r.Run(":" + port)
}
