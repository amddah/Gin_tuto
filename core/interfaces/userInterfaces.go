package interfaces

import (
	"gin_api/core/models"

	"github.com/google/uuid"
)

type UserService interface {
	Get(uid uuid.UUID) (*models.User, error)
	GetAll() ([]models.User, error)
	GetByID(id string) (*models.User, error)
}

type UserRepository interface {
	FindByID(uid uuid.UUID) (*models.User, error)
	FindAll() ([]models.User, error)
	FindByStringID(id string) (*models.User, error)
}
