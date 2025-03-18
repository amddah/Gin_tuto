package interfaces

import (
	"gin_api/core/readmodels"
)

// PostReadRepository for accessing the post read model (MongoDB)
type PostReadRepository interface {
	FindByID(id string) (*readmodels.PostReadModel, error)
	FindAll() ([]readmodels.PostReadModel, error)
	Save(post *readmodels.PostReadModel) error
}

// PostReadService for querying post data
type PostReadService interface {
	GetByID(id string) (*readmodels.PostReadModel, error)
	GetAll() ([]readmodels.PostReadModel, error)
}
