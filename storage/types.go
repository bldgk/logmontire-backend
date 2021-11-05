package storage

import (
	"logmontire/common/errors"
)

type Savable interface {
	GetInstance() interface{}
	GetId() string
}

type Repository interface {
	GetCollectionName() string
	Save(Savable) ([]interface{}, *errors.Error)
}
