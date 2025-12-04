package service

import (
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/repository"
)

type IUserService interface {
	Insert(user *model.User) error
}

type UserServiceImpl struct {
	repo repository.IUserRepository
}

func NewIUserService(repo repository.IUserRepository) IUserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (service UserServiceImpl) Insert(user *model.User) error {
	return service.repo.Insert(user)
}
