package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/xcodeme21/go-crm-service/routes"
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
            "message": "Yayyyy I'm Gin Gonic",
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
    db, err := database.Initialize()
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }
    r.Use(database.Inject(db))

    // Seeder
    database.VoucherCategories()

    routes.VoucherCategoriesRoutes(r)
    routes.VouchersRoutes(r)

    port := os.Getenv("PORT")
    r.NoRoute(lostInSpce)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Error starting the server: %v", err)
    }
}
