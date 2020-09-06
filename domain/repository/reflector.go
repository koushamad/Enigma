package repository

import "github.com/koushamad/Enigma/domain/entity"

type Reflector interface {
	GetAll() []entity.Reflect
	GetByIndex(index int) (entity.Reflect, error)
	Add(r *entity.Reflect) error
}
