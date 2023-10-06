package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xcodeme21/go-crm-service/models"
	"gorm.io/gorm"
)

type VoucherCategoriesController struct {
	DB    *gorm.DB
	DBTwo *gorm.DB
}

func (c *VoucherCategoriesController) ListCategories(ctx *gin.Context) {
	var products []models.VoucherCategories
	err := c.DB.Order("id asc").Find(&products).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"data": products, "error_message": err.Error(), "status": ctx.Writer.Status()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": products, "error_message": nil, "status": ctx.Writer.Status()})
}