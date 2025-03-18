package service

import (
	"gin_api/core/interfaces"
	"gin_api/core/models"

	"github.com/google/uuid"
)

// Implémentation de UserService
type UserServiceImpl struct {
	UserRepo interfaces.UserRepository
}

// Constructeur pour UserServiceImpl
func NewUserService(userRepo interfaces.UserRepository) interfaces.UserService {
	return &UserServiceImpl{UserRepo: userRepo}
}

// Méthode pour récupérer un utilisateur via le repository
func (s *UserServiceImpl) Get(uid uuid.UUID) (*models.User, error) {
	return s.UserRepo.FindByID(uid)
}

func (s *UserServiceImpl) GetAll() ([]models.User, error) {
	return s.UserRepo.FindAll()
}

func (s *UserServiceImpl) GetByID(id string) (*models.User, error) {
	return s.UserRepo.FindByStringID(id)
}
