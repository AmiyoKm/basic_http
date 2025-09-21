package user

import (
	"fmt"

	"github.com/AmiyoKm/basic_http/domain"
)

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Create(user *domain.Users) (*domain.Users, error) {
	u, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}
func (s *service) Delete(id string) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) GetByEmail(email string) (*domain.Users, error) {
	u, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, err
	}
	return u, nil
}
func (s *service) GetByID(id string) (*domain.Users, error) {
	u, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (s *service) Update(user *domain.Users) (*domain.Users, error) {

	u, err := s.userRepo.GetByID(user.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
