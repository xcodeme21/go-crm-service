package providers

import (
	"github.com/xcodeme21/go-crm-service/models"
	"gorm.io/gorm"
	"fmt"
	utils "github.com/xcodeme21/go-crm-service/utils"
)

type VouchersProvider interface {
	FindAll(filters models.FilterVouchers) ([]models.Vouchers, int64)
	GetCategoryByID(id int) (*models.Vouchers, error)
	Create(newCategory models.Vouchers) (*models.Vouchers, error)
	Update(id int, updatedCategory models.Vouchers) (*models.Vouchers, error)
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

func (b *DBVouchersProvider) FindAll(filters models.FilterVouchers) ([]models.Vouchers, int64) {
	var data []models.Vouchers
	var dataCount []models.Vouchers

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
		q.Where("voucher_name ILIKE ?", "%"+filters.Search+"%")
		q.Or("series_id = ?", filters.Search)
	}

	q.Find(&dataCount).Count(&count)

	// Untuk Export tidak perlu pagination
	if filters.Page == -1 && filters.PerPage == -1 {
		q.Find(&data)
	} else {
		q.Scopes(utils.Paginate(filters.Page, filters.PerPage)).Find(&data)
	}
	//fmt.Println(data)
	return data, count
}

func (p *DBVouchersProvider) GetCategoryByID(id int) (*models.Vouchers, error) {
	// Implementation to retrieve a category by ID from the database
	var category models.Vouchers
	err := p.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (p *DBVouchersProvider) Create(newCategory models.Vouchers) (*models.Vouchers, error) {
	// Insert the new category into the database
	if err := p.DB.Create(&newCategory).Error; err != nil {
		return nil, err
	}
	return &newCategory, nil
}

func (p *DBVouchersProvider) Update(id int, updatedCategory models.Vouchers) (*models.Vouchers, error) {
	// Fetch the existing category by ID
	var existingCategory models.Vouchers
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
	return p.DB.Delete(&models.Vouchers{}, id).Error
}
