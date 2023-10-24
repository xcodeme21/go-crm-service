package controllers

import (
	"cloud.google.com/go/storage"


	"net/http"
	"strconv"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"os"
	"io"
	"math/big"
	"net/url"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xcodeme21/go-crm-service/api/cms/vouchers/services"
	"github.com/xcodeme21/go-crm-service/models"
	"github.com/xcodeme21/go-crm-service/utils"
	"google.golang.org/appengine"
	"google.golang.org/api/option"
	"encoding/json"
	"time"
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
	var newCategory models.Voucher
	fmt.Println(newCategory)

	// Bind the request body to the newCategory struct
	if err := ctx.ShouldBindJSON(&newCategory); err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	fmt.Println("tes2")
	var credentials models.GoogleCloudCredential

	storageBucket := os.Getenv("GOOGLE_CLOUD_STORAGE_BUCKET")
	appengineCtx := appengine.NewContext(ctx.Request) // Renamed the context variable

	credentials = helper.GetGoogleCloudStorageCredentials()
	jsonCredential, err := json.Marshal(credentials)

	if err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	storageClient, err = storage.NewClient(appengineCtx, option.WithCredentialsJSON(jsonCredential))
	if err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	var source = "crm_vouchers"

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	size := int(file.(Sizer).Size())
	max := 1024 * 1024 * 2

	if size > max {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, "File yang Anda upload lebih besar dari 2 Mb")
		return
	}

	defer file.Close()

	fileName := header.Filename
	arrFileName := strings.Split(fileName, ".")
	ext := arrFileName[len(arrFileName)-1]

	newTime := time.Now().Format(helper.GetDateFormat("YYYYMMDDHHIISS"))

	bi := big.NewInt(0)
	h := md5.New()
	h.Write([]byte(newTime))
	hexstr := hex.EncodeToString(h.Sum(nil))
	bi.SetString(hexstr, 16)

	fileNameMD5 := bi.String()

	sw := storageClient.Bucket(storageBucket).Object(source + "/" + fileNameMD5 + "." + ext).NewWriter(appengineCtx)

	if _, err := io.Copy(sw, file); err != nil {
		helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	if err := sw.Close(); err != nil {
		helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	u, err := url.Parse("/" + storageBucket + "/" + sw.Attrs().Name)
	if err != nil {
		helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	newCategory.Image = "https://storage.googleapis.com" + u.EscapedPath() // Corrected field name from image to Image

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
