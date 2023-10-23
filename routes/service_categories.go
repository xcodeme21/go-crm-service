package routes

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/xcodeme21/go-crm-service/api/cms/voucher_categories/controllers"
	"github.com/xcodeme21/go-crm-service/api/cms/voucher_categories/providers"
	"github.com/xcodeme21/go-crm-service/api/cms/voucher_categories/services"
	"github.com/xcodeme21/go-crm-service/database"
)

func VoucherCategoriesRoutes(r *gin.Engine) {
	db, err := database.Connect() // Call the Connect function to get the database connection
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    voucherCategoriesService := services.NewVoucherCategoriesService(providers.NewDBVoucherCategoriesProvider(db))

	// Dereference the pointer to pass the underlying value to the constructor
	categoriesController := controllers.NewVoucherCategoriesController(*voucherCategoriesService)

	cmsGroup := r.Group("/api/cms/")
	cmsGroup.GET("/voucher-categories", categoriesController.ListCategories)
	cmsGroup.GET("/voucher-categories/:id", categoriesController.DetailCategory)
	cmsGroup.POST("/voucher-categories", categoriesController.CreateCategory)
	cmsGroup.PUT("/voucher-categories/:id", categoriesController.UpdateCategory)
	cmsGroup.DELETE("/voucher-categories/:id", categoriesController.DeleteCategory)
}