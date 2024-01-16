package controllers

import (
	"cloud.google.com/go/storage"
	//"fmt"

	"net/http"
	"strconv"
	"strings"
	// "crypto/md5"
	// "encoding/hex"
	// "os"
	// "io"
	// "math/big"
	// "net/url"
	// "google.golang.org/appengine"
	// "google.golang.org/api/option"
	// "encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xcodeme21/go-crm-service/api/cms/vouchers/services"
	"github.com/xcodeme21/go-crm-service/models"
	"github.com/xcodeme21/go-crm-service/utils"
)

var (
	storageClient *storage.Client
)

type Sizer interface {
	Size() int64
}

type VouchersController struct {
	service services.VouchersService
}

func NewVouchersController(service services.VouchersService) *VouchersController {
	return &VouchersController{
		service: service,
	}
}

func (c *VouchersController) FindAll(ctx *gin.Context) {
	var paginate models.Pagination
	
	start := ctx.Query("start_date")
	end := ctx.Query("end_date")
	status := ctx.Query("status")
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))

	filters := models.FilterVoucher{
		Start:   start,
		End:     end,
		Status:  status,
		Search:  search,
		Page:    page,
		PerPage: perPage,
		SortBy:  ctx.DefaultQuery("sort_by", "id"),
		SortDir: strings.ToUpper(ctx.DefaultQuery("sort_dir", "DESC")),
	}

	otherFilters := map[string]string{
	}

	data, count := c.service.FindAll(filters)

	paginate = helper.GenerateBasicPagination(ctx, otherFilters, data, count)

	helper.JSONResponse(ctx, http.StatusOK, paginate, "")
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
	category, err := c.service.Detail(id)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusNotFound, nil, "Category not found")
		return
	}

	helper.JSONResponse(ctx, http.StatusOK, category, "")
}

func (c *VouchersController) Create(ctx *gin.Context) {
    var voucherRequest models.VoucherRequest

    err := ctx.ShouldBind(&voucherRequest)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	_, err = c.service.Detail(voucherRequest.CategoryId)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusNotFound, nil, "Category not found")
		return
	}
    
    // Create a Voucher object
    result := models.Voucher{
        VoucherName:       voucherRequest.VoucherName,
        SeriesId:          voucherRequest.SeriesId,
        IsExternalVoucher: voucherRequest.IsExternalVoucher,
        CampaignId:        voucherRequest.CampaignId,
        CampaignVoucherId: voucherRequest.CampaignVoucherId,
        Point:             voucherRequest.Point,
        Tier:              voucherRequest.Tier,
        CategoryId:        voucherRequest.CategoryId,
        IsLimited:         voucherRequest.IsLimited,
        StartDate:         voucherRequest.StartDate,
        EndDate:           voucherRequest.EndDate,
        Image:             voucherRequest.Image,
        Description:       voucherRequest.Description,
        Status:            voucherRequest.Status,
        CreatedAt:         time.Now(),
        UpdatedAt:         time.Now(),
    }

    // Save the Voucher to the database
    createdCategory, err := c.service.Create(result)
    if err != nil {
        // Handle database save error
        helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
        return
    }

    helper.JSONResponse(ctx, http.StatusOK, createdCategory, "Voucher created successfully")
}

func (c *VouchersController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, "Invalid category ID")
		return
	}

	var updatedCategory models.Voucher

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
	_, err = c.service.Detail(id)
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
