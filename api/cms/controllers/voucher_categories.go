package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xcodeme21/go-crm-service/api/cms/services"
	"github.com/xcodeme21/go-crm-service/utils"
	"github.com/xcodeme21/go-crm-service/models"
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

func (c *VoucherCategoriesController) CreateCategory(ctx *gin.Context) {
    var newCategory models.VoucherCategories

    // Bind the request body to the newCategory struct
    if err := ctx.ShouldBindJSON(&newCategory); err != nil {
        helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
        return
    }

    createdCategory, err := c.service.CreateCategory(newCategory)
    if err != nil {
        helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
        return
    }

    helper.JSONResponse(ctx, http.StatusCreated, createdCategory, "")
}

func (c *VoucherCategoriesController) UpdateCategory(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        helper.JSONResponse(ctx, http.StatusBadRequest, nil, "Invalid category ID")
        return
    }

    var updatedCategory models.VoucherCategories

    if err := ctx.ShouldBindJSON(&updatedCategory); err != nil {
        helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
        return
    }

    // Call the service to update the category by ID
    updatedCategoryPtr, err := c.service.UpdateCategory(id, updatedCategory)
    if err != nil {
        helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
        return
    }

    // Convert the pointer to a regular struct, if needed
    updatedCategory = *updatedCategoryPtr

    helper.JSONResponse(ctx, http.StatusOK, updatedCategory, "")
}


func (c *VoucherCategoriesController) DeleteCategory(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        helper.JSONResponse(ctx, http.StatusBadRequest, nil, "Invalid category ID")
        return
    }

    err = c.service.DeleteCategory(id)
    if err != nil {
        helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
        return
    }

    helper.JSONResponse(ctx, http.StatusOK, nil, "Category deleted successfully")
}

