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
