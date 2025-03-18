package service

import (
	"gin_api/core/interfaces"
	"gin_api/core/readmodels"

	"github.com/google/uuid"
)

// UserReadServiceImpl implements the read service for users
type UserReadServiceImpl struct {
	UserReadRepo interfaces.UserReadRepository
}

// NewUserReadService creates a new user read service
func NewUserReadService(userReadRepo interfaces.UserReadRepository) interfaces.UserReadService {
	return &UserReadServiceImpl{UserReadRepo: userReadRepo}
}

// Get retrieves a user by UUID
func (s *UserReadServiceImpl) Get(uid uuid.UUID) (*readmodels.UserReadModel, error) {
	return s.UserReadRepo.FindByUUID(uid)
}

// GetAll retrieves all users
func (s *UserReadServiceImpl) GetAll() ([]readmodels.UserReadModel, error) {
	return s.UserReadRepo.FindAll()
}

// GetByID retrieves a user by string ID
func (s *UserReadServiceImpl) GetByID(id string) (*readmodels.UserReadModel, error) {
	return s.UserReadRepo.FindByID(id)
}
