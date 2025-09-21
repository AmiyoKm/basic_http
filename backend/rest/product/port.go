package product

import "github.com/AmiyoKm/basic_http/domain"

type Service interface {
	Get() ([]*domain.Product, error)
	GetByID(id string) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(ID string) error
}
