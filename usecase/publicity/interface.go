package publicity

import (
	"github.com/louisbillaut/test/entity"
	"time"
)

//Writer publicity interface
type Writer interface {
	Create(pub *entity.Publicity) (entity.ID, error)
}

//Reader publicity interface
type Reader interface {
	GetByTimeStamp(from time.Time, to time.Time) ([]*entity.Publicity, error)
	GetByOS(os string) ([]*entity.Publicity, error)
	GetByID(id entity.ID) (*entity.Publicity, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetPublicity(id entity.ID) (*entity.Publicity, error)
	GetByTimeStamp(from time.Time, to time.Time) ([]*entity.Publicity, error)
	GetByOS(os string) ([]*entity.Publicity, error)
	CreatePublicity(types string, ua string, ip string) (entity.ID, error)
}