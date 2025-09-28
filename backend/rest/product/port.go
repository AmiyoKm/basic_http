package product

import "github.com/AmiyoKm/basic_http/domain"

type Service interface {
	Get(page, limit int) ([]*domain.Product, error)
	Count() (int, error)
	GetByID(id string) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(ID string) error
}
