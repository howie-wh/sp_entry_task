package dao

import (
	"entry_task/pkg/dao"
	"entry_task/pkg/logger"
	"entry_task/pkg/models/response"
	"entry_task/pkg/table"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

type SysUserDao struct {
}

//查询公共sql
func (d SysUserDao) querySql() *xorm.Session {
	return dao.SqlDB.NewSession().Table("sys_user_tab")
}

// GetSysUserById 根据id查询用户数据
func (d SysUserDao) GetSysUserById(userId uint64) *response.SysUserResponse {
	var resp response.SysUserResponse
	get, err := d.querySql().Where("user_id = ?", userId).Get(&resp)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil
	}
	if !get {
		return nil
	}
	return &resp
}

// GetSysUserByUserName 根据用户名查询用户数据
func (d SysUserDao) GetSysUserByUserName(user table.SysUser) *table.SysUser {
	i, err := dao.SqlDB.Get(&user)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil
	}
	if i {
		return &user
	}
	return nil
}