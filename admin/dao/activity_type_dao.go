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

type ActivityTypeDao struct {
}

//查询公共sql
func (d ActivityTypeDao) querySql() *xorm.Session {
	return dao.SqlDB.NewSession().Table("activity_type_tab")
}

// FindList 查询集合
func (d ActivityTypeDao) FindList(query req.ActivityTypeListQuery) ([]*response.ActivityTypeResponse, int64) {
	resp := make([]*response.ActivityTypeResponse, 0)
	sql := d.querySql()
	if !gotool.StrUtils.HasEmpty(query.TypeName) {
		sql.And("type_name like concat('%',?,'%')", query.TypeName)
	}
	sql.And("del_flag = 0")

	total, _ := page.GetTotal(sql.Clone())
	err := sql.Limit(query.Offset, query.Start).Find(&resp)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil, 0
	}
	return resp, total
}

// GetActivityTypeById 根据id查询数据
func (d ActivityTypeDao) GetActivityTypeById(typeId uint64) *response.ActivityTypeResponse {
	var resp response.ActivityTypeResponse
	get, err := d.querySql().Where("type_id = ?", typeId).Get(&resp)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil
	}
	if !get {
		return nil
	}
	return &resp
}

// GetActivityTypeByTypeName 根据类型名称查询活动类型
func (d ActivityTypeDao) GetActivityTypeByTypeName(activityType table.ActivityType) *table.ActivityType {
	i, err := dao.SqlDB.Get(&activityType)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil
	}
	if i {
		return &activityType
	}
	return nil
}

// Insert 添加活动类型
func (d ActivityTypeDao) Insert(body req.ActivityTypeBody) *table.ActivityType {
	tab := &table.ActivityType{
		TypeName: body.TypeName,
		DelFlag: "0",
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession().Table("activity_type_tab")
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
func (d ActivityTypeDao) Update(body req.ActivityTypeBody) int64 {
	tab := &table.ActivityType{
		TypeName: body.TypeName,
		DelFlag: body.DelFlag,
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession().Table("activity_type_tab")
	session.Begin()
	_, err := session.Where("type_id = ?", body.TypeId).Update(tab)
	if err != nil {
		session.Rollback()
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return 0
	}
	session.Commit()
	return 1
}

// Remove 根据id删除数据
//
func (d ActivityTypeDao) Remove(typeIds []uint64) int64 {
	tab := &table.ActivityType{
		DelFlag: "1",
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession().Table("activity_type_tab")
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
