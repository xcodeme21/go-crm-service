package services

import (
	"github.com/xcodeme21/go-crm-service/api/cms/vouchers/providers"
	"github.com/xcodeme21/go-crm-service/models"
)

type VouchersService struct {
	provider providers.VouchersProvider
}

func NewVouchersService(provider providers.VouchersProvider) *VouchersService {
	return &VouchersService{
		provider: provider,
	}
}

func (s *VouchersService) ListCategories() ([]models.Vouchers, error) {
	return s.provider.ListCategories()
}

func (s *VouchersService) GetCategoryByID(id int) (*models.Vouchers, error) {
	category, err := s.provider.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *VouchersService) CreateCategory(newCategory models.Vouchers) (*models.Vouchers, error) {
	// Validate newCategory and perform any necessary business logic
	// Create the new category in the provider
	createdCategory, err := s.provider.CreateCategory(newCategory)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

func (s *VouchersService) UpdateCategory(id int, updatedCategory models.Vouchers) (*models.Vouchers, error) {
	// Validate updatedCategory and perform any necessary business logic
	// Update the category in the provider
	updatedCategoryFromProvider, err := s.provider.UpdateCategory(id, updatedCategory)
	if err != nil {
		return nil, err
	}
	return updatedCategoryFromProvider, nil
}

func (s *VouchersService) DeleteCategory(id int) error {
	// Perform any necessary checks before deleting the category
	return s.provider.DeleteCategory(id)
}
