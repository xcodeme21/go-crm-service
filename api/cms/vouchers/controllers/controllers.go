package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xcodeme21/go-crm-service/api/cms/vouchers/services"
	"github.com/xcodeme21/go-crm-service/models"
	helper "github.com/xcodeme21/go-crm-service/utils"
)

type VouchersController struct {
	service services.VouchersService
}

func NewVouchersController(service services.VouchersService) *VouchersController {
	return &VouchersController{
		service: service,
	}
}

func (c *VouchersController) ListVouchers(ctx *gin.Context) {
	products, err := c.service.ListVouchers()
	if err != nil {
		helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	helper.JSONResponse(ctx, http.StatusOK, products, "")
}

func (c *VouchersController) Detail(ctx *gin.Context) {
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

func (c *VouchersController) Create(ctx *gin.Context) {
	var newCategory models.Vouchers

	// Bind the request body to the newCategory struct
	if err := ctx.ShouldBindJSON(&newCategory); err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	createdCategory, err := c.service.Create(newCategory)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	helper.JSONResponse(ctx, http.StatusCreated, createdCategory, "")
}

func (c *VouchersController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, "Invalid category ID")
		return
	}

	var updatedCategory models.Vouchers

	if err := ctx.ShouldBindJSON(&updatedCategory); err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	// Call the service to update the category by ID
	updatedCategoryPtr, err := c.service.Update(id, updatedCategory)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// Convert the pointer to a regular struct, if needed
	updatedCategory = *updatedCategoryPtr

	helper.JSONResponse(ctx, http.StatusOK, updatedCategory, "")
}

func (c *VouchersController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, "Invalid category ID")
		return
	}

	// Check if the category exists
	_, err = c.service.GetCategoryByID(id)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusNotFound, nil, "Category not found")
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	helper.JSONResponse(ctx, http.StatusOK, nil, "Category deleted successfully")
}
