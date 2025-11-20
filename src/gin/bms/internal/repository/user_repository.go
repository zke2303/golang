package repository

import (
	"errors"

	"github.com/zhang/bms/internal/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindById(id uint64) (*model.User, error)
	Create(user *model.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repo UserRepositoryImpl) FindById(id uint64) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo UserRepositoryImpl) Create(user *model.User) error {
	return repo.db.Create(user).Error
}
