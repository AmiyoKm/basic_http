package product

import (
	"github.com/AmiyoKm/basic_http/domain"
	"github.com/AmiyoKm/basic_http/rest/product"
)

type Service interface {
	product.Service
}

type ProductRepo interface {
	Get(page, limit int) ([]*domain.Product, error)
	Count() (int, error)
	GetByID(id string) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(ID string) error
}
