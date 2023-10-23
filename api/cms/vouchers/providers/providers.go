package providers

import (
	"github.com/xcodeme21/go-crm-service/models"
	"gorm.io/gorm"
)

type VouchersProvider interface {
	ListCategories() ([]models.Vouchers, error)
	GetCategoryByID(id int) (*models.Vouchers, error)
	CreateCategory(newCategory models.Vouchers) (*models.Vouchers, error)
	UpdateCategory(id int, updatedCategory models.Vouchers) (*models.Vouchers, error)
	DeleteCategory(id int) error
}

type DBVouchersProvider struct {
	DB *gorm.DB
}

func NewDBVouchersProvider(DB *gorm.DB) VouchersProvider {
	return &DBVouchersProvider{
		DB: DB,
	}
}

func (p *DBVouchersProvider) ListCategories() ([]models.Vouchers, error) {
	var categories []models.Vouchers
	err := p.DB.Order("id asc").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
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

func (p *DBVouchersProvider) CreateCategory(newCategory models.Vouchers) (*models.Vouchers, error) {
	// Insert the new category into the database
	if err := p.DB.Create(&newCategory).Error; err != nil {
		return nil, err
	}
	return &newCategory, nil
}

func (p *DBVouchersProvider) UpdateCategory(id int, updatedCategory models.Vouchers) (*models.Vouchers, error) {
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

func (p *DBVouchersProvider) DeleteCategory(id int) error {
	// Delete the category by ID
	return p.DB.Delete(&models.Vouchers{}, id).Error
}
