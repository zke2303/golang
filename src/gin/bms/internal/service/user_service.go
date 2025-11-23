package service

import (
	"github.com/zhang/bms/internal/dto"
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/repository"
	"github.com/zhang/bms/internal/utils"
)

type IUserService interface {
	FindById(id uint64) (*model.User, error)
	Create(dto *dto.UserRequest) error
	Delete(id uint64) error
	Update(id uint64, req *dto.UserUpdateRequest) error
	Login(login *dto.LoginDTO) (*string, error)
}

type UserServiceImpl struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) Login(login *dto.LoginDTO) (*string, error) {
	err := s.repo.Login(login)
	if err != nil {
		return nil, err
	}
	// 生成 token
	token, err := utils.GeneratJwt(login.Username)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *UserServiceImpl) FindById(id uint64) (*model.User, error) {
	return s.repo.FindById(id)
}

func (s *UserServiceImpl) Create(dto *dto.UserRequest) error {
	// 将 UserRequest 对象转换成 User对象
	user := dto.ToModel()
	return s.repo.Create(user)
}

func (s *UserServiceImpl) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *UserServiceImpl) Update(id uint64, req *dto.UserUpdateRequest) error {
	// 1.将UserUpdateRequest对象转换成Map对象
	updateData := req.ToMap()
	// 2.特殊处理
	// 如果用户请求中包含密码,进行加密
	if req.Password != nil && *req.Password != "" {
		// 暂时忽略
		updateData["password"] = req.Password
	}

	// 3.判断 map 是否为空, 如果为空,则证明没有字段需要更新,直接返回
	if len(updateData) == 0 {
		return nil
	}

	// 4.调用repository
	return s.repo.Update(id, updateData)
}
