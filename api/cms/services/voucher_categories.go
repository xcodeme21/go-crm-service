package services

import (
	"github.com/xcodeme21/go-crm-service/models"
	"github.com/xcodeme21/go-crm-service/api/cms/providers"
)

type VoucherCategoriesService struct {
	provider providers.VoucherCategoriesProvider
}

func NewVoucherCategoriesService(provider providers.VoucherCategoriesProvider) *VoucherCategoriesService {
	return &VoucherCategoriesService{
		provider: provider,
	}
}

func (s *VoucherCategoriesService) ListCategories() ([]models.VoucherCategories, error) {
	return s.provider.ListCategories()
}

func (s *VoucherCategoriesService) GetCategoryByID(id int) (*models.VoucherCategories, error) {
    category, err := s.provider.GetCategoryByID(id)
    if err != nil {
        return nil, err
    }
    return category, nil
}

func (s *VoucherCategoriesService) CreateCategory(newCategory models.VoucherCategories) (*models.VoucherCategories, error) {
    // Validate newCategory and perform any necessary business logic
    // Create the new category in the provider
    createdCategory, err := s.provider.CreateCategory(newCategory)
    if err != nil {
        return nil, err
    }
    return createdCategory, nil
}

func (s *VoucherCategoriesService) UpdateCategory(id int, updatedCategory models.VoucherCategories) (*models.VoucherCategories, error) {
    // Validate updatedCategory and perform any necessary business logic
    // Update the category in the provider
    updatedCategoryFromProvider, err := s.provider.UpdateCategory(id, updatedCategory)
    if err != nil {
        return nil, err
    }
    return updatedCategoryFromProvider, nil
}


func (s *VoucherCategoriesService) DeleteCategory(id int) error {
    // Perform any necessary checks before deleting the category
    return s.provider.DeleteCategory(id)
}

