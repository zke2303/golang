package repository

import (
	"errors"

	"github.com/mattn/go-sqlite3"
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/pkg/errcode"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Insert(user *model.User) error
	Delete(id uint64) error
	FindByUsername(username string) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewIUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

// FindByUsername
// 根据用户名查询用户信息
func (repo UserRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.NotFound
		}
		return nil, errcode.InternalServerError
	}

	return &user, nil
}
func (repo UserRepositoryImpl) Insert(user *model.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errcode.Exists
		}
		// 2) SQLite 专用判断（这是你缺少的部分）
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.Code == sqlite3.ErrConstraint {
			return errcode.Exists
		}
		return err
	}
	return nil
}

func (repo UserRepositoryImpl) Delete(id uint64) error {
	result := repo.db.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errcode.NotFound
	}

	return nil
}
