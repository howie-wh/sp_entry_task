package dao

import (
	"entry_task/pkg/dao"
	"entry_task/pkg/logger"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
	"entry_task/pkg/page"
	"entry_task/pkg/table"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

type UserDao struct {
}

func (d UserDao) querySql() *xorm.Session {
	return dao.SqlDB.NewSession().Table("user_tab")
}

// FindList 查询用户集合
func (d UserDao) FindList(query req.UserListQuery) ([]*response.UserResponse, int64) {
	resp := make([]*response.UserResponse, 0)
	sql := d.querySql()
	sql.And("del_flag = 0")

	total, _ := page.GetTotal(sql.Clone())
	err := sql.Limit(query.Offset, query.Start).Find(&resp)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil, 0
	}
	return resp, total
}

// GetUserById 根据id查询用户数据
func (d UserDao) GetUserById(userId int64) *response.UserResponse {
	var resp response.UserResponse
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

// GetUserByUserName 根据用户名查询用户数据
func (d UserDao) GetUserByUserName(user table.User) *table.User {
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