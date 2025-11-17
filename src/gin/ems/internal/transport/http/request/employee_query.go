package request

import (
	"github.com/zhang/bms/internal/model"
)

type EmployeeQuery struct {
	model.Employee
	Gender *int `form:"gender" json:"gender"`
	Offset int  `form:"offset,default=1" json:"offset"`
	Limit  int  `form:"limit,default=10" json:"limit"`
}
