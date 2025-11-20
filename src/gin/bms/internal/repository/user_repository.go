package repository

import (
	"errors"

	"github.com/zhang/bms/internal/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindById(id uint64) (*model.User, error)
	Create(dto *model.User) error
	Delete(id uint64) error
	Update(id uint64, data map[string]interface{}) error
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

func (repo UserRepositoryImpl) Delete(id uint64) error {
	err := repo.db.Delete(&model.User{}, id).Error
	return err
}

func (repo UserRepositoryImpl) Update(id uint64, data map[string]interface{}) error {
	result := repo.db.Model(&model.User{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
