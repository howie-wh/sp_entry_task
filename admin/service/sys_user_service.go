package service

import (
	"entry_task/admin/dao"
	"entry_task/pkg/models/response"
	"entry_task/pkg/table"
)

// SysUserService 用户操作业务逻辑
type SysUserService struct {
	sysUserDao dao.SysUserDao
}

// GetSysUserById 根据id查询用户数据
func (s SysUserService) GetSysUserById(userId uint64) *response.SysUserResponse {
	return s.sysUserDao.GetSysUserById(userId)
}

// GetSysUserByUserName 根据用户名查询用户
func (s SysUserService) GetSysUserByUserName(name string) *table.SysUser {
	user := table.SysUser{}
	user.UserName = name
	return s.sysUserDao.GetSysUserByUserName(user)
}