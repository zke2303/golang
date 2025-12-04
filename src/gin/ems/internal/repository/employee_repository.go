package repository

import (
	"fmt"

	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/transport/http/request"
	"github.com/zhang/bms/internal/transport/http/response"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindById(id uint64) (model.Employee, error)
	List(query request.EmployeeQuery) (response.PageResult[model.Employee], error)
	Insert(e *model.Employee) error
	Delete(id uint64) error
	Update(m *model.Employee) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (r *employeeRepository) FindById(id uint64) (model.Employee, error) {
	var employee model.Employee
	err := r.db.First(&employee, id).Error
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (r *employeeRepository) List(q request.EmployeeQuery) (response.PageResult[model.Employee], error) {
	var list []model.Employee
	query := r.db.Model(&model.Employee{})

	if q.Username != "" {
		query = query.Where("username like ?", "%"+q.Username+"%")
	}

	if q.Nickname != "" {
		query = query.Where("nickname like ?", "%"+q.Nickname+"%")
	}

	if q.Department != "" {
		query = query.Where("department = ?", q.Department)
	}

	if q.Gender != nil {
		query = query.Where("gender = ?", q.Gender)
	}

	// 统计总数
	var total int64
	query.Count(&total)

	// --- 分页条件 (Limit and Offset) ---

	// 1. 设置每页大小 (Limit)
	// 通常会对 Limit 进行校验，确保它在一个合理的范围内 (例如 1 到 100)
	if q.Limit > 0 {
		query = query.Limit(q.Limit)
	} else {
		// 设置默认 Limit，例如 20
		query = query.Limit(20)
	}

	// 2. 设置偏移量 (Offset)
	if q.Offset > 0 {
		query = query.Offset(q.Offset)
	}

	if err := query.Find(&list).Error; err != nil {
		return response.PageResult[model.Employee]{}, err
	}
	for _, value := range list {
		fmt.Println(value)
	}

	// 封装成 PageResult对象
	var page response.PageResult[model.Employee]
	page.Total = total
	page.Current = q.Offset
	page.Records = list

	return page, nil
}

func (r *employeeRepository) Insert(e *model.Employee) error {
	return r.db.Create(e).Error
}

func (r *employeeRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Employee{}, id).Error
}

func (r *employeeRepository) Update(m *model.Employee) error {
	return r.db.Model(&model.Employee{}).Where("id = ?", m.Id).Updates(m).Error
}
