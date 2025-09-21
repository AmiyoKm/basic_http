package user

import "github.com/AmiyoKm/basic_http/domain"

type Service interface {
	GetByID(id string) (*domain.Users, error)
	GetByEmail(email string) (*domain.Users, error)
	Create(user *domain.Users) (*domain.Users, error)
	Update(user *domain.Users) (*domain.Users, error)
	Delete(id string) error
}
