package repository

import (
	"gin_api/core/interfaces"
	"gin_api/core/models"
	"gin_api/initializer"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// Constructeur pour UserRepositoryImpl
func NewUserRepository() interfaces.UserRepository {
	return &UserRepositoryImpl{DB: initializer.DB}
}

// Méthode pour récupérer un utilisateur par ID
func (repo *UserRepositoryImpl) FindByID(uid uuid.UUID) (*models.User, error) {
	var user models.User
	if err := repo.DB.First(&user, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepositoryImpl) FindByStringID(id string) (*models.User, error) {
	var user models.User
	if err := repo.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
