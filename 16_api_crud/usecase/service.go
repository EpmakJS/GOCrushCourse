package usecase

import "16_api_crud/entities"

// Service interface
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// GetAll returns list of products
func (s *Service) GetAll() ([]*entities.Product, error) {
	return s.repo.List()
}
