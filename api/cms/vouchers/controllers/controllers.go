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

// func (c *VouchersController) Creates(ctx *gin.Context) {
//     var voucherRequest models.VoucherRequest

//     // Bind the request body to the voucherRequest struct
//     if err := ctx.ShouldBindJSON(&voucherRequest); err != nil {
//         // Handle JSON binding error
//         helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
//         return
//     }

//     var credentials models.GoogleCloudCredential

//     // Retrieve Google Cloud Storage bucket name from environment variable
//     storageBucket := os.Getenv("GOOGLE_CLOUD_STORAGE_BUCKET")
//     appengineCtx := appengine.NewContext(ctx.Request)

//     // Get Google Cloud Storage credentials and create a storage client
//     credentials = helper.GetGoogleCloudStorageCredentials()
//     jsonCredential, err := json.Marshal(credentials)

//     if err != nil {
//         // Handle credential error
//         helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
//         return
//     }

//     storageClient, err := storage.NewClient(appengineCtx, option.WithCredentialsJSON(jsonCredential))
//     if err != nil {
//         // Handle storage client creation error
//         helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
//         return
//     }

//     // Define the source and validate and process the uploaded file
//     var source = "crm_vouchers"
//     file, err := ctx.FormFile("file")
//     if err != nil {
//         // Handle file retrieval error
//         helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
//         return
//     }
// 	// Open the uploaded file
// file, err := file.Open()
// if err != nil {
//     // Handle file opening error
//     helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
//     return
// }
// defer file.Close() // Close the file when you're done with it

// // Validate the file size
// size := int(file.(Sizer).Size())
// max := 1024 * 1024 * 2

// if size > max {
//     // Handle file size validation error
//     helper.JSONResponse(ctx, http.StatusBadRequest, nil, "File size exceeds the maximum allowed.")
//     return
// }


//     defer file.Close()

//     // Generate a unique filename for the uploaded file
//     fileName := file.Filename
//     arrFileName := strings.Split(fileName, ".")
//     ext := arrFileName[len(arrFileName)-1]

//     newTime := time.Now().Format(helper.GetDateFormat("YYYYMMDDHHIISS"))

//     bi := big.NewInt(0)
//     h := md5.New()
//     h.Write([]byte(newTime))
//     hexstr := hex.EncodeToString(h.Sum(nil))
//     bi.SetString(hexstr, 16)

//     fileNameMD5 := bi.String()

//     // Create a writer to upload the file to Google Cloud Storage
//     sw := storageClient.Bucket(storageBucket).Object(source + "/" + fileNameMD5 + "." + ext).NewWriter(appengineCtx)

//     if _, err := io.Copy(sw, file); err != nil {
//         // Handle file upload error
//         helper.JSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
//         return
//     }

//     if err := sw.Close(); err != nil {
//         // Handle writer closing error
//         helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
//         return
//     }

//     // Construct the URL of the uploaded file
//     u, err := url.Parse("/" + storageBucket + "/" + sw.Attrs().Name)
//     if err != nil {
//         // Handle URL parsing error
//         helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
//         return
//     }

//     voucherRequest.Image = "https://storage.googleapis.com" + u.EscapedPath()

//     // Create a Voucher object
//     result := models.Voucher{
//         VoucherName:       voucherRequest.VoucherName,
//         SeriesId:          voucherRequest.SeriesId,
//         IsExternalVoucher: voucherRequest.IsExternalVoucher,
//         CampaignId:        voucherRequest.CampaignId,
//         CampaignVoucherId: voucherRequest.CampaignVoucherId,
//         Point:             voucherRequest.Point,
//         Tier:              voucherRequest.Tier,
//         CategoryId:        voucherRequest.CategoryId,
//         IsLimited:         voucherRequest.IsLimited,
//         StartDate:         voucherRequest.StartDate,
//         EndDate:           voucherRequest.EndDate,
//         Image:             voucherRequest.Image,
//         Description:       voucherRequest.Description,
//         Status:            voucherRequest.Status,
//         CreatedAt:         time.Now(),
//         UpdatedAt:         time.Now(),
//     }

//     // Save the Voucher to the database
//     createdCategory, err := c.service.Create(result)
//     if err != nil {
//         // Handle database save error
//         helper.JSONResponse(ctx, http.StatusInternalServerError, nil, err.Error())
//         return
//     }

//     // Respond with the created Voucher
//     helper.JSONResponse(ctx, http.StatusCreated, createdCategory, "")
// }



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
