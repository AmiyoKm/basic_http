package product

import (
	"github.com/AmiyoKm/basic_http/domain"
	"github.com/AmiyoKm/basic_http/rest/product"
)

type Service interface {
	product.Service
}

type ProductRepo interface {
	Get() ([]*domain.Product, error)
	GetByID(id string) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(ID string) error
}
