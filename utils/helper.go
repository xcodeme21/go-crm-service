package helper

import (
	"errors"
	"fmt"
	"math"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/xcodeme21/go-crm-service/models"
)

func JSONResponse(ctx *gin.Context, status int, data interface{}, errorMessage string) {
	ctx.JSON(status, gin.H{
		"status":        status,
		"data":          data,
		"error_message": errorMessage,
	})
}

func GenerateBasicPagination(c *gin.Context, otherFilters map[string]string, items interface{}, count int64) (pagination models.Pagination) {
	var paginate models.Pagination
	var filters models.FiltersPagination
	var strField string
	var strValue string
	routeURL := c.Request.URL.Path
	params := url.Values{}

	filters.Search = c.Query("search")
	filters.Page = c.DefaultQuery("page", "1")
	filters.PerPage = c.DefaultQuery("per_page", "10")

	filters.SortBy = c.DefaultQuery("sort_by", "id")
	filters.SortDir = strings.ToUpper(c.DefaultQuery("sort_dir", "DESC"))

	params.Add("page", "x") // Will be replace with pagination page link
	params.Add("per_page", filters.PerPage)
	params.Add("sort_by", filters.SortBy)
	params.Add("sort_dir", filters.SortDir)

	// TODO DYNAMIC FILTER. OTHER DYNAMIC FILTER ONLY STRING TYPE.
	for field, def := range otherFilters {
		strField = CamelCase(field)
		if IsExistField(filters, strField) {
			strValue = c.DefaultQuery(field, def)
			_ = SetDynamicStringField(&filters, strField, strValue)
			if strValue != "" {
				params.Add(field, strValue)
			}
		}
	}

	paginate.Total = count
	page, _ := strconv.Atoi(filters.Page)
	perPage, _ := strconv.Atoi(filters.PerPage)
	paginate.CurrentPage = page
	paginate.PerPage = perPage
	totalPage := int(math.Ceil(float64(count) / float64(perPage)))

	paginate.Filters = filters

	if paginate.Total == 0 {
		paginate.Items = nil
	} else {
		paginate.Items = items

		paginate.LastPage = totalPage
		strURL := routeURL + "?" + params.Encode()

		paginate.FirstPageURL = strings.Replace(strURL, "page=x", "page=1", 1)
		if page > 1 && page <= totalPage {
			paginate.PrevPageURL = strings.Replace(strURL, "page=x", "page="+strconv.Itoa(page-1), 1)
		}
		if page > 0 && page < totalPage {
			paginate.NextPageURL = strings.Replace(strURL, "page=x", "page="+strconv.Itoa(page+1), 1)
		}
		paginate.LastPageURL = strings.Replace(strURL, "page=x", "page="+strconv.Itoa(totalPage), 1)

		from := ((page - 1) * perPage) + 1
		to := int64(from + perPage - 1)
		paginate.From = int64(from)
		paginate.To = to
		if to > count {
			paginate.To = count
		}
	}

	return paginate

}

func IsExistField(i interface{}, name string) bool {
	ci := reflect.ValueOf(i)
	fi := ci.FieldByName(name)
	if !fi.IsValid() {
		return false
	}
	return true
}

func CamelCase(str string) string {
	var match = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")
	return match.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

func SetDynamicStringField(v interface{}, name string, value string) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("v must be pointer to struct")
	}
	rv = rv.Elem()
	fv := rv.FieldByName(name)
	/*if !fv.IsValid() {
		return fmt.Errorf("not a field name: %s", name)
	}*/
	if !fv.CanSet() {
		return fmt.Errorf("cannot set field %s", name)
	}
	if fv.Kind() != reflect.String {
		return fmt.Errorf("%s is not a string field", name)
	}
	fv.SetString(value)
	return nil
}

func Paginate(page int, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case perPage > 100:
			perPage = 100
		case perPage <= 0:
			perPage = 10
		}

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}

func Tiering(tier string) string {
	tierArray := strings.Split(tier, ",")
	var response []string

	for _, item := range tierArray {
		var rs = ""

		if item == "1" {
			rs = "Crew"
		} else if item == "2" {
			rs = "Co-Pilot"
		} else if item == "3" {
			rs = "Pilot"
		} else {
			rs = "Spacetronot"
		}

		response = append(response, rs)
	}

	// Join the tiers with a comma separator
	tierString := strings.Join(response, ",")
	fmt.Println(tierString)

	return tierString
}

