package interfaces

import (
	"gin_api/core/readmodels"

	"github.com/google/uuid"
)

// UserReadRepository for accessing the read model (MongoDB)
type UserReadRepository interface {
	FindByID(id string) (*readmodels.UserReadModel, error)
	FindAll() ([]readmodels.UserReadModel, error)
	Save(user *readmodels.UserReadModel) error
	FindByUUID(uid uuid.UUID) (*readmodels.UserReadModel, error)
}

// UserReadService for querying user data
type UserReadService interface {
	GetByID(id string) (*readmodels.UserReadModel, error)
	GetAll() ([]readmodels.UserReadModel, error)
	Get(uid uuid.UUID) (*readmodels.UserReadModel, error)
}
