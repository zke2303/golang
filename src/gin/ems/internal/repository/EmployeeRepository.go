package repository

import (
	"fmt"

	"github.com/zhangke/ems/internal/model"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindByID(id uint) (*model.Employee, error)
	Create(m *model.Employee) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) FindByID(id uint) (*model.Employee, error) {
	var emp model.Employee
	if err := r.db.First(&emp, id).Error; err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *employeeRepository) Create(e *model.Employee) error {
	err := r.db.Create(e).Error
	if err != nil {
		fmt.Println("数据库插入数据失败")
		return err
	}
	return nil
}
