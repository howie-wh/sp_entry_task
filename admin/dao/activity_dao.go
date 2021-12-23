package dao

import (
	"entry_task/pkg/dao"
	"entry_task/pkg/logger"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
	"entry_task/pkg/page"
	"entry_task/pkg/table"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"time"
)

type ActivityDao struct {
}

//查询公共sql
func (d ActivityDao) querySql() *xorm.Session {
	return dao.SqlDB.NewSession().Table("activity_tab")
}

// FindList 查询活动集合
func (d ActivityDao) FindList(query req.ActivityListQuery) ([]*response.ActivityResponse, int64) {
	respList := make([]*response.ActivityResponse, 0)
	tabList := make([]*table.Activity, 0)

	sql := d.querySql()
	sql.Cols("activity_id, type_id, title, location, content, start_time, end_time")
	if query.TypeId > 0 {
		sql.And("type_id = ? ", query.TypeId)
	}
	if query.StartTime > 0 {
		sql.And("start_time > ?", query.StartTime)
	}
	if query.EndTime > 0 {
		sql.And("end_time < ?", query.EndTime)
	}
	sql.And("del_flag = 0")
	total, _ := page.GetTotal(sql.Clone())
	gotool.Logs.ErrorLog().Println(total)
	err := sql.Limit(query.Offset, query.Start).Find(&tabList)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil, 0
	}

	for _, tab := range tabList {
		resp := &response.ActivityResponse {
			ActivityId: tab.ActivityId,
			TypeId: tab.TypeId,
			Title: tab.Title,
			Location: tab.Location,
			Content: tab.Content,
			StartTime: tab.StartTime,
			EndTime: tab.EndTime,
		}
		respList = append(respList, resp)
	}
	return respList, total
}

// GetActivityById 根据id查询数据
func (d ActivityDao) GetActivityById(activityId uint64) *response.ActivityResponse {
	var resp response.ActivityResponse
	get, err := d.querySql().Where("activity_id = ?", activityId).Get(&resp)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil
	}
	if !get {
		return nil
	}
	return &resp
}

// GetActivityByTitle 根据用户名查询用户数据
func (d ActivityDao) GetActivityByTitle(activity table.Activity) *table.Activity {
	i, err := dao.SqlDB.Get(&activity)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil
	}
	if i {
		return &activity
	}
	return nil
}

// Insert 添加活动
func (d ActivityDao) Insert(body req.ActivityBody) *table.Activity {
	tab := &table.Activity{
		TypeId: body.TypeId,
		Title: body.Title,
		Location: body.Location,
		Content: body.Content,
		StartTime: body.StartTime,
		EndTime: body.EndTime,
		DelFlag: "0",
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession().Table("activity_tab")
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
func (d ActivityDao) Update(body req.ActivityBody) int64 {
	tab := &table.Activity{
		TypeId: body.TypeId,
		Title: body.Title,
		Location: body.Location,
		Content: body.Content,
		StartTime: body.StartTime,
		EndTime: body.EndTime,
		DelFlag: body.DelFlag,
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession().Table("activity_tab")
	session.Begin()
	_, err := session.Where("activity_id = ?", body.ActivityId).Update(tab)
	if err != nil {
		session.Rollback()
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return 0
	}
	session.Commit()
	return 1
}

// Remove 根据id删除数据
func (d ActivityDao) Remove(activityIds []uint64) int64 {
	tab := &table.Activity{
		DelFlag: "1",
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession().Table("activity_tab")
	session.Begin()
	_, err := session.In("activity_id", activityIds).Update(tab)
	if err != nil {
		session.Rollback()
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return 0
	}
	session.Commit()
	return 1
}

// RemoveByTypeIds 根据id删除数据
func (d ActivityDao) RemoveByTypeIds(typeIds []uint64) int64 {
	tab := &table.Activity{
		DelFlag: "1",
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession().Table("activity_tab")
	session.Begin()
	_, err := session.In("type_id", typeIds).Update(tab)
	if err != nil {
		session.Rollback()
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return 0
	}
	session.Commit()
	return 1
}