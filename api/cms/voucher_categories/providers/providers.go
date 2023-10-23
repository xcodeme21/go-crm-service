package providers

import (
	"github.com/xcodeme21/go-crm-service/models"
	"gorm.io/gorm"
)

type VoucherCategoriesProvider interface {
	FindAll() ([]models.VoucherCategories, error)
	Detail(id int) (*models.VoucherCategories, error)
	Create(newCategory models.VoucherCategories) (*models.VoucherCategories, error)
	Update(id int, updatedCategory models.VoucherCategories) (*models.VoucherCategories, error)
	Delete(id int) error
}

type DBVoucherCategoriesProvider struct {
	DB *gorm.DB
}

func NewDBVoucherCategoriesProvider(DB *gorm.DB) VoucherCategoriesProvider {
	return &DBVoucherCategoriesProvider{
		DB: DB,
	}
}

func (p *DBVoucherCategoriesProvider) FindAll() ([]models.VoucherCategories, error) {
	var categories []models.VoucherCategories
	err := p.DB.Order("id asc").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (p *DBVoucherCategoriesProvider) Detail(id int) (*models.VoucherCategories, error) {
	// Implementation to retrieve a category by ID from the database
	var category models.VoucherCategories
	err := p.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (p *DBVoucherCategoriesProvider) Create(newCategory models.VoucherCategories) (*models.VoucherCategories, error) {
	// Insert the new category into the database
	if err := p.DB.Create(&newCategory).Error; err != nil {
		return nil, err
	}
	return &newCategory, nil
}

func (p *DBVoucherCategoriesProvider) Update(id int, updatedCategory models.VoucherCategories) (*models.VoucherCategories, error) {
	// Fetch the existing category by ID
	var existingCategory models.VoucherCategories
	if err := p.DB.First(&existingCategory, id).Error; err != nil {
		return nil, err
	}

	// Update the existing category with the new data
	if err := p.DB.Model(&existingCategory).Updates(updatedCategory).Error; err != nil {
		return nil, err
	}

	return &existingCategory, nil
}

func (p *DBVoucherCategoriesProvider) Delete(id int) error {
	// Delete the category by ID
	return p.DB.Delete(&models.VoucherCategories{}, id).Error
}
