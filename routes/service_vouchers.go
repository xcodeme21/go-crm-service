package routes

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/xcodeme21/go-crm-service/api/cms/vouchers/controllers"
	"github.com/xcodeme21/go-crm-service/api/cms/vouchers/providers"
	"github.com/xcodeme21/go-crm-service/api/cms/vouchers/services"
	"github.com/xcodeme21/go-crm-service/database"
)

func VouchersRoutes(r *gin.Engine) {
	db, err := database.Connect() // Call the Connect function to get the database connection
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    VouchersService := services.NewVouchersService(providers.NewDBVouchersProvider(db))

	// Dereference the pointer to pass the underlying value to the constructor
	categoriesController := controllers.NewVouchersController(*VouchersService)

	cmsGroup := r.Group("/api/cms/")
	cmsGroup.GET("/vouchers", categoriesController.ListCategories)
	cmsGroup.GET("/vouchers/:id", categoriesController.DetailCategory)
	cmsGroup.POST("/vouchers", categoriesController.CreateCategory)
	cmsGroup.PUT("/vouchers/:id", categoriesController.UpdateCategory)
	cmsGroup.DELETE("/vouchers/:id", categoriesController.DeleteCategory)
}