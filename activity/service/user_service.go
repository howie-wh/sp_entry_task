package service

import (
	"entry_task/activity/dao"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
	"entry_task/pkg/table"
)

// UserService 用户操作业务逻辑
type UserService struct {
	userDao dao.UserDao
}

// FindList 查询用户集合业务方法
func (s UserService) FindList(query *req.UserListQuery) ([]*response.UserResponse, int64) {
	return s.userDao.FindList(query)
}

// GetUserById 根据id查询用户数据
func (s UserService) GetUserById(userId uint64) *response.UserResponse {
	return s.userDao.GetUserById(userId)
}

// GetUserByUserName 根据用户名查询用户
func (s UserService) GetUserByUserName(name string) *table.User {
	user := &table.User{}
	user.UserName = name
	return s.userDao.GetUserByUserName(user)
}

// Insert 添加用户业务逻辑
func (s UserService) Insert(body *req.UserBody) bool {
	user := s.userDao.InsertUser(body)
	if user != nil {
		return true
	}
	return false
}
