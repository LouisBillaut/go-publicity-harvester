package publicity

import (
	"github.com/louisbillaut/test/entity"
	"time"
)

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

//CreatePublicity Create a publicity
func (s *Service) CreatePublicity(pubType string, ua string, ip string) (entity.ID, error) {
	e := entity.NewPublicity(pubType, ua, ip)
	return s.repo.Create(e)
}

//GetByOS Gets by OS
func (s *Service) GetByOS(os string) ([]*entity.Publicity, error) {
	return s.repo.GetByOS(os)
}

//GetByTimeStamp Gets by OS
func (s *Service) GetByTimeStamp(from time.Time, to time.Time) ([]*entity.Publicity, error) {
	return s.repo.GetByTimeStamp(from, to)
}

//GetPublicity Gets by publicity by ID
func (s *Service) GetPublicity(id entity.ID) (*entity.Publicity, error) {
	return s.repo.GetByID(id)
}