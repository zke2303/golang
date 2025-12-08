package service

import (
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/repository"
)

type IUserService interface {
	Insert(user *model.User) error
	Delete(id uint64) error
	FindByUsername(username string) (*model.User, error)
}

type UserServiceImpl struct {
	repo repository.IUserRepository
}

func NewIUserService(repo repository.IUserRepository) IUserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

// FindByUsername
func (service UserServiceImpl) FindByUsername(username string) (*model.User, error) {
	return service.repo.FindByUsername(username)
}

func (service UserServiceImpl) Insert(user *model.User) error {
	return service.repo.Insert(user)
}

func (service UserServiceImpl) Delete(id uint64) error {
	return service.repo.Delete(id)
}
