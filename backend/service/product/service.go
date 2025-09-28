package product

import "github.com/AmiyoKm/basic_http/domain"

type service struct {
	repo ProductRepo
}

func NewService(repo ProductRepo) Service {
	return &service{
		repo: repo,
	}
}
func (s *service) Get(page, limit int) ([]*domain.Product, error) {
	return s.repo.Get(page, limit)
}

func (s *service) Count() (int, error) {
	return s.repo.Count()
}

func (s *service) Create(product *domain.Product) error {
	err := s.repo.Create(product)
	return err
}

func (s *service) Delete(ID string) error {
	return s.repo.Delete(ID)
}

func (s *service) GetByID(id string) (*domain.Product, error) {
	return s.repo.GetByID(id)
}

func (s *service) Update(product *domain.Product) error {
	err := s.repo.Update(product)
	return err
}
