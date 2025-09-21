package user

import (
	"github.com/AmiyoKm/basic_http/domain"
	userHandler "github.com/AmiyoKm/basic_http/rest/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(user *domain.Users) (*domain.Users, error)
	Delete(id string) error
	GetByEmail(email string) (*domain.Users, error)
	GetByID(id string) (*domain.Users, error)
	Update(user *domain.Users) (*domain.Users, error)
}
