package publicity_type

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

func (s *Service) GetPublicityTypeByPublicityID(publicityID entity.ID) (*entity.PublicityType, error) {
	return s.repo.GetByPublicityID(publicityID)
}