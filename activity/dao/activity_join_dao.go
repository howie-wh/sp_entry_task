package dao

import (
	"entry_task/pkg/dao"
	"entry_task/pkg/logger"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
	"entry_task/pkg/table"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"time"
)

type ActivityJoinDao struct {
}

//查询公共sql
func (d ActivityJoinDao) querySql() *xorm.Session {
	return dao.SqlDB.NewSession().Table("activity_join_tab")
}

// GetActivityJoinById 根据id查询数据
func (d ActivityJoinDao) GetActivityJoinById(activityId, userId uint64) *response.ActivityJoinResponse {
	var resp response.ActivityJoinResponse
	sql := d.querySql()
	get, err := sql.Where("activity_id = ? and user_id = ?", activityId, userId).Get(&resp)
	sql.And("del_flag = '0'")
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil
	}
	if !get {
		return nil
	}
	return &resp
}


// Insert 添加活动报名
func (d ActivityJoinDao) Insert(body *req.ActivityJoinBody) *table.ActivityJoin {
	tab := &table.ActivityJoin{
		ActivityId: body.ActivityId,
		UserId: body.UserId,
		DelFlag: "0",
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession().Table("activity_join_tab")
	session.Begin()
	_, err := session.Insert(tab)
	if err != nil {
		session.Rollback()
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil
	}
	session.Commit()
	return tab
}

// Update 修改数据
func (d ActivityJoinDao) Update(body *req.ActivityJoinBody) int64 {
	tab := &table.ActivityJoin{
		DelFlag: body.DelFlag,
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession().Table("activity_join_tab")
	session.Begin()
	_, err := session.Where("activity_id = ? and user_id = ?", body.ActivityId, body.UserId).Update(tab)
	if err != nil {
		session.Rollback()
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return 0
	}
	session.Commit()
	return 1
}

// Remove 根据id删除数据
func (d ActivityJoinDao) Remove(activityId, userId uint64) int64 {
	body := &req.ActivityJoinBody{
		ActivityId: activityId,
		UserId: userId,
		DelFlag: "1",
	}
	return d.Update(body)
}
