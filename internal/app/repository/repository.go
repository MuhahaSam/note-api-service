package repository

import "github.com/google/uuid"

type Repository interface {
	Create(object any) (uuid.UUID, error)
	Read(Id uuid.UUID) (any error)
	Update(Id uuid.UUID, object any) error
	Delete(Id uuid.UUID) error
}
