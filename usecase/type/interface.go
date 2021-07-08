package _type

import (
	"github.com/louisbillaut/test/entity"
)

//Writer type interface
type Writer interface {
	Create(pub *entity.Type) (entity.ID, error)
}

//Reader type interface
type Reader interface {
	GetByName(name string) (*entity.Type, error)
	GetByID(id entity.ID) (*entity.Type, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	CreateType(pub *entity.Type) (entity.ID, error)
	GetTypeByName(name string) (*entity.Type, error)
	GetTypeByID(id entity.ID) (*entity.Type, error)
}