package publicity_type

import (
	"github.com/louisbillaut/test/entity"
)

//Writer publicityType interface
type Writer interface {
	//TODO
}

//Reader publicityType interface
type Reader interface {
	GetByPublicityID(id entity.ID) (*entity.PublicityType, error)
	//TODO GetByTypeID
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetPublicityTypeByPublicityID(publicityID entity.ID) (*entity.PublicityType, error)
}
