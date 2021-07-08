package _type

import "github.com/louisbillaut/test/entity"

//Service interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateType Create a type
func (s *Service) CreateType(typ string) (entity.ID, error) {
	e, err := entity.NewType(typ)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

//GetTypeByName Gets a type
func (s *Service) GetTypeByName(name string) (*entity.Type, error) {
	return s.repo.GetByName(name)
}

//GetTypeByID Gets a type
func (s *Service) GetTypeByID(id entity.ID) (*entity.Type, error) {
	return s.repo.GetByID(id)
}
