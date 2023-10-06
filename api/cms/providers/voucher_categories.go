package providers

import (
	"github.com/xcodeme21/go-crm-service/models"
	"gorm.io/gorm"
)

type VoucherCategoriesProvider interface {
	ListCategories() ([]models.VoucherCategories, error)
    GetCategoryByID(id int) (*models.VoucherCategories, error)
}

type DBVoucherCategoriesProvider struct {
	DB *gorm.DB
}

func NewDBVoucherCategoriesProvider(DB *gorm.DB) VoucherCategoriesProvider {
	return &DBVoucherCategoriesProvider{
		DB: DB,
	}
}

func (p *DBVoucherCategoriesProvider) ListCategories() ([]models.VoucherCategories, error) {
	var categories []models.VoucherCategories
	err := p.DB.Order("id asc").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (p *DBVoucherCategoriesProvider) GetCategoryByID(id int) (*models.VoucherCategories, error) {
    // Implementation to retrieve a category by ID from the database
    var category models.VoucherCategories
    err := p.DB.First(&category, id).Error
    if err != nil {
        return nil, err
    }
    return &category, nil
}