package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhangke/ems/internal/model"
	"github.com/zhangke/ems/internal/repository"
)

type EmployeeService interface {
	FindById(c *gin.Context) (*model.Employee, error)
	Create(c *gin.Context) error
}

type employeeService struct {
	Repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{Repo: repo}
}

func (s *employeeService) FindById(c *gin.Context) (*model.Employee, error) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	return s.Repo.FindByID(uint(id))
}

func (s *employeeService) Create(c *gin.Context) error {
	var employee model.Employee
	err := c.ShouldBind(&employee)
	if err != nil {
		fmt.Println("数据格式转换错误")
		return err
	}

	err = s.Repo.Create(&employee)
	if err != nil {
		fmt.Printf("插入数据失败")
		return err
	}

	return nil
}
