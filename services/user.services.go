package services

import (
	"github.com/Dom-HTG/gin/models"
	"github.com/Dom-HTG/gin/repository"
)

type UserServiceContainer interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(id int, user *models.User)
	DeleteUser(id int) error
}

type UserServiceDependency struct {
	repo repository.UserRepositoryContainer
}

func NewUserServiceDependency(repo repository.UserRepositoryContainer) *UserServiceDependency {
	return &UserServiceDependency{
		repo: repo,
	}
}

func (s *UserServiceDependency) CreateUser(user *models.User) error {
	err := s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceDependency) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceDependency) UpdateUser(id int, user *models.User) error {
	err := s.repo.UpdateUser(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceDependency) DeleteUser(id int) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
