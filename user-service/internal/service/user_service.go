package service

import (
	"user-service/internal/model"
	"user-service/internal/repository"
)

type UserService interface {
	RegisterUser(username string) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	ListUsers() ([]model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) RegisterUser(username string) (*model.User, error) {
	user := &model.User{Username: username}
	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(id int) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) GetUserByUsername(username string) (*model.User, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *userService) ListUsers() ([]model.User, error) {
	return s.repo.ListUsers()
}
