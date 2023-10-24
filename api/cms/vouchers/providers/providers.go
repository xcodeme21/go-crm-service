package providers

import (
	"github.com/xcodeme21/go-crm-service/models"
	"gorm.io/gorm"
	"fmt"
	utils "github.com/xcodeme21/go-crm-service/utils"
)

type VouchersProvider interface {
	FindAll(filters models.FilterVoucher) ([]models.VoucherFindAllResponse, int64)
	Detail(id int) (*models.Voucher, error)
	Create(newCategory models.Voucher) (*models.Voucher, error)
	Update(id int, updatedCategory models.Voucher) (*models.Voucher, error)
	Delete(id int) error
}

type DBVouchersProvider struct {
	DB *gorm.DB
}

func NewDBVouchersProvider(DB *gorm.DB) VouchersProvider {
	return &DBVouchersProvider{
		DB: DB,
	}
}

func (b *DBVouchersProvider) FindAll(filters models.FilterVoucher) ([]models.VoucherFindAllResponse, int64) {
	var data []models.Voucher
	var dataCount []models.Voucher
	var response []models.VoucherFindAllResponse

	var count int64
	sortBy := "id"
	sortDir := "DESC"

	if filters.SortBy != "" {
		sortBy = filters.SortBy
	}
	if filters.SortDir == "ASC" {
		sortDir = filters.SortDir
	}
	q := b.DB.Order(fmt.Sprintf(`%s %s`, sortBy, sortDir))

	if filters.Status != "" {
		q.Where("status = ?", filters.Status)
	}

	if filters.Start != "" && filters.End == "" {
		q.Where("start_date >= ?", filters.Start)
	} else if filters.Start == "" && filters.End != "" {
		q.Where("end_date <= ?", filters.End)
	} else if filters.Start != "" && filters.End != "" {
		q.Where("start_date >= ? AND end_date <= ?", filters.Start, filters.End)
	}	

	if filters.Search != "" {
		q.Where("series_id = ?", filters.Search)
	}
	

	q.Find(&dataCount).Count(&count)

	// Untuk Export tidak perlu pagination
	if filters.Page == -1 && filters.PerPage == -1 {
		q.Find(&data)
	} else {
		q.Scopes(utils.Paginate(filters.Page, filters.PerPage)).Find(&data)
	}

	for _, item := range data {
		var responseSingle models.VoucherFindAllResponse

        category := b.FindCategoryRelationByCategoryID(item.CategoryId)

        responseSingle.ID = item.ID
        responseSingle.VoucherName = item.VoucherName
        responseSingle.SeriesId = item.SeriesId
        responseSingle.IsExternalVoucher = item.IsExternalVoucher
        responseSingle.CampaignId = item.CampaignId
        responseSingle.CampaignVoucherId = item.CampaignVoucherId
        responseSingle.Point = item.Point
        responseSingle.Tier = item.Tier
        responseSingle.TierName = utils.Tiering(item.Tier)
        responseSingle.CategoryId = item.CategoryId
        responseSingle.CategoryName = category.Name
        responseSingle.IsLimited = item.IsLimited
        responseSingle.StartDate = item.StartDate
        responseSingle.EndDate = item.EndDate
        responseSingle.Image = item.Image
        responseSingle.Description = item.Description
        responseSingle.Status = item.Status
        responseSingle.CreatedAt = item.CreatedAt
        responseSingle.UpdatedAt = item.UpdatedAt

        response = append(response, responseSingle) // Tambahkan item ke slice dataCount (atau slice lainnya yang Anda buat)
    }
	
	//fmt.Println(data)
	return response, count
}

func (b *DBVouchersProvider) FindCategoryRelationByCategoryID(id int) models.VoucherCategory {
	var rs models.VoucherCategory
	b.DB.Where("id = ?", id).First(&rs)

	return rs
}

func (p *DBVouchersProvider) Detail(id int) (*models.Voucher, error) {
	// Implementation to retrieve a category by ID from the database
	var category models.Voucher
	err := p.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (p *DBVouchersProvider) Create(newCategory models.Voucher) (*models.Voucher, error) {
	// Insert the new category into the database
	if err := p.DB.Create(&newCategory).Error; err != nil {
		return nil, err
	}
	return &newCategory, nil
}

func (p *DBVouchersProvider) Update(id int, updatedCategory models.Voucher) (*models.Voucher, error) {
	// Fetch the existing category by ID
	var existingCategory models.Voucher
	if err := p.DB.First(&existingCategory, id).Error; err != nil {
		return nil, err
	}

	// Update the existing category with the new data
	if err := p.DB.Model(&existingCategory).Updates(updatedCategory).Error; err != nil {
		return nil, err
	}

	return &existingCategory, nil
}

func (p *DBVouchersProvider) Delete(id int) error {
	// Delete the category by ID
	return p.DB.Delete(&models.Voucher{}, id).Error
}
