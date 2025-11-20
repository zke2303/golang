package service

import (
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/repository"
)

type IUserService interface {
	FindById(id uint64) (*model.User, error)
	Create(m *model.User) error
}

type UserServiceImpl struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) FindById(id uint64) (*model.User, error) {
	return s.repo.FindById(id)
}

func (s *UserServiceImpl) Create(user *model.User) error {
	return s.repo.Create(user)
}
