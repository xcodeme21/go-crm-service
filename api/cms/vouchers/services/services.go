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

func (s *VouchersService) FindAll(filters models.FilterVouchers) ([]models.VouchersFindAllResponse, int64) {
	return s.provider.FindAll(filters)
}

func (s *VouchersService) Detail(id int) (*models.Vouchers, error) {
	category, err := s.provider.Detail(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *VouchersService) Create(newCategory models.Vouchers) (*models.Vouchers, error) {
	// Validate newCategory and perform any necessary business logic
	// Create the new category in the provider
	createdCategory, err := s.provider.Create(newCategory)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

func (s *VouchersService) Update(id int, updatedCategory models.Vouchers) (*models.Vouchers, error) {
	// Validate updatedCategory and perform any necessary business logic
	// Update the category in the provider
	updatedCategoryFromProvider, err := s.provider.Update(id, updatedCategory)
	if err != nil {
		return nil, err
	}
	return updatedCategoryFromProvider, nil
}

func (s *VouchersService) Delete(id int) error {
	// Perform any necessary checks before deleting the category
	return s.provider.Delete(id)
}
