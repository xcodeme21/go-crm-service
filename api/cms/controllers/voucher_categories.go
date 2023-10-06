package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xcodeme21/go-crm-service/api/cms/services"
	"github.com/xcodeme21/go-crm-service/utils"
)

type VoucherCategoriesController struct {
	service services.VoucherCategoriesService
}

func NewVoucherCategoriesController(service services.VoucherCategoriesService) *VoucherCategoriesController {
	return &VoucherCategoriesController{
		service: service,
	}
}

func (c *VoucherCategoriesController) ListCategories(ctx *gin.Context) {
	products, err := c.service.ListCategories()
	if err != nil {
		helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	helper.JSONResponse(ctx, http.StatusOK, products, "")
}



func (c *VoucherCategoriesController) DetailCategory(ctx *gin.Context) {
    // Get the category ID from the route parameter
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        helper.JSONResponse(ctx, http.StatusBadRequest, nil, "Invalid category ID")
        return
    }

    // Call the service to retrieve the category details by ID
    category, err := c.service.GetCategoryByID(id)
    if err != nil {
        helper.JSONResponse(ctx, http.StatusNotFound, nil, "Category not found")
        return
    }

    helper.JSONResponse(ctx, http.StatusOK, category, "")
}
