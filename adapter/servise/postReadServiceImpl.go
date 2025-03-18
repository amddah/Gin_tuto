package service

import (
	"gin_api/core/interfaces"
	"gin_api/core/readmodels"
)

// PostReadServiceImpl implements the post read service
type PostReadServiceImpl struct {
	PostReadRepo interfaces.PostReadRepository
}

// NewPostReadService creates a new post read service
func NewPostReadService(postReadRepo interfaces.PostReadRepository) interfaces.PostReadService {
	return &PostReadServiceImpl{PostReadRepo: postReadRepo}
}

// GetByID retrieves a post by ID
func (s *PostReadServiceImpl) GetByID(id string) (*readmodels.PostReadModel, error) {
	return s.PostReadRepo.FindByID(id)
}

// GetAll retrieves all posts
func (s *PostReadServiceImpl) GetAll() ([]readmodels.PostReadModel, error) {
	return s.PostReadRepo.FindAll()
}
