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
)

type ActivityDao struct {
}

//查询公共sql
func (d ActivityDao) querySql() *xorm.Session {
	return dao.SqlDB.NewSession().Table([]string{"activity_tab", "a"})
}

// FindList 查询活动集合
func (d ActivityDao) FindList(query *req.ActivityListQuery) ([]*response.ActivityResponse, int64) {
	respList := make([]*response.ActivityResponse, 0)
	tabList := make([]*table.Activity, 0)

	sql := d.querySql()
	if query.TypeId > 0 {
		sql.And("a.type_id = ? ", query.TypeId)
	}
	if query.StartTime > 0 {
		sql.And("a.start_time > ?", query.StartTime)
	}
	if query.EndTime > 0 {
		sql.And("a.end_time < ?", query.EndTime)
	}
	sql.And("a.del_flag = '0'")

	total, _ := page.GetTotal(sql.Clone())
	err := sql.Asc("a.start_time").Limit(query.Offset, query.Start).Find(&tabList)
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
			IsJoin: false,
		}
		respList = append(respList, resp)
	}
	return respList, total
}

// GetActivityById 根据id查询数据
func (d ActivityDao) GetActivityById(query *req.ActivityQuery) *response.ActivityResponse {
	var resp response.ActivityResponse
	sql := d.querySql()
	if query.UserId > 0 {
		sql.Select("a.activity_id, a.type_id, a.title, a.location, a.content, a.start_time, a.end_time, if(isNULL(j.join_id), '0','1') as is_join").
			Join("LEFT", []string{"activity_join_tab", "j"},
				"a.activity_id in (select j.activity_id from activity_join_tab j where j.del_flag = '0' and j.user_id = ?)",query.UserId)
	} else {
		sql.Select("a.activity_id, a.type_id, a.title, a.location, a.content, a.start_time, a.end_time, 0 as is_join")
	}

	sql.And("a.activity_id = ? ", query.ActivityId)
	sql.And("a.del_flag = 0")
	get, err := sql.Get(&resp)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil
	}
	if !get {
		return nil
	}
	return &resp
}

// FindUserActivityList 查询活动集合
func (d ActivityDao) FindUserActivityList(query *req.ActivityListQuery) ([]*response.ActivityResponse, int64) {
	respList := make([]*response.ActivityResponse, 0)
	tabList := make([]*response.ActivityTempResponse, 0)

	sql := dao.SqlDB.NewSession()
	total, _ := page.GetTotal(d.querySql().Clone())
	err := sql.Sql("SELECT a.activity_id, a.type_id, a.title, a.location, a.content, a.start_time, a.end_time, j.join_id " +
		"FROM activity_tab AS a " +
		"LEFT JOIN (select join_id, activity_id from activity_join_tab where activity_join_tab.user_id = 102 and activity_join_tab.del_flag = '0') as j " +
		"on a.activity_id = j.activity_id").Find(&tabList)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil, 0
	}
	for _, tab := range tabList {
		gotool.Logs.ErrorLog().Println(tab.JoinId)
		resp := &response.ActivityResponse {
			ActivityId: tab.ActivityId,
			TypeId: tab.TypeId,
			Title: tab.Title,
			Location: tab.Location,
			Content: tab.Content,
			StartTime: tab.StartTime,
			EndTime: tab.EndTime,
		}
		if tab.JoinId == 0 {
			resp.IsJoin = false
		} else {
			resp.IsJoin = true
		}
		respList = append(respList, resp)
	}

	return respList, total
}