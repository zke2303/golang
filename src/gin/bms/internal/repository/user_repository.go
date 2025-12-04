package repository

import (
	"errors"

	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/pkg/errcode"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Insert(user *model.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewIUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repo UserRepositoryImpl) Insert(user *model.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errcode.Exists
		}
		return err
	}
	return nil
}
