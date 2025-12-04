package service

import (
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/repository"
	"github.com/zhang/bms/internal/transport/http/request"
	"github.com/zhang/bms/internal/transport/http/response"
)

type EmployeeService struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) EmployeeService {
	return EmployeeService{repository: repository}
}

func (s EmployeeService) FindById(id uint64) (model.Employee, error) {
	return s.repository.FindById(id)
}

func (s EmployeeService) List(query request.EmployeeQuery) (response.PageResult[model.Employee], error) {
	return s.repository.List(query)
}

// Insert 往 employees 表中插入一条数据/*
func (s EmployeeService) Insert(m *model.Employee) error {
	return s.repository.Insert(m)
}

func (s EmployeeService) Delete(id uint64) error {
	return s.repository.Delete(id)
}

func (s EmployeeService) Update(m *model.Employee) error {
	return s.repository.Update(m)
}
