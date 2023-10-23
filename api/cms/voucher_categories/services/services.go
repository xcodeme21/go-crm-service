package services

import (
	"github.com/xcodeme21/go-crm-service/api/cms/voucher_categories/providers"
	"github.com/xcodeme21/go-crm-service/models"
)

type VoucherCategoriesService struct {
	provider providers.VoucherCategoriesProvider
}

func NewVoucherCategoriesService(provider providers.VoucherCategoriesProvider) *VoucherCategoriesService {
	return &VoucherCategoriesService{
		provider: provider,
	}
}

func (s *VoucherCategoriesService) FindAll() ([]models.VoucherCategories, error) {
	return s.provider.FindAll()
}

func (s *VoucherCategoriesService) Detail(id int) (*models.VoucherCategories, error) {
	category, err := s.provider.Detail(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *VoucherCategoriesService) Create(newCategory models.VoucherCategories) (*models.VoucherCategories, error) {
	// Validate newCategory and perform any necessary business logic
	// Create the new category in the provider
	createdCategory, err := s.provider.Create(newCategory)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

func (s *VoucherCategoriesService) Update(id int, updatedCategory models.VoucherCategories) (*models.VoucherCategories, error) {
	// Validate updatedCategory and perform any necessary business logic
	// Update the category in the provider
	updatedCategoryFromProvider, err := s.provider.Update(id, updatedCategory)
	if err != nil {
		return nil, err
	}
	return updatedCategoryFromProvider, nil
}

func (s *VoucherCategoriesService) Delete(id int) error {
	// Perform any necessary checks before deleting the category
	return s.provider.Delete(id)
}
