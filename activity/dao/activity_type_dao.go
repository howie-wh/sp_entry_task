package dao

import (
	"entry_task/pkg/dao"
	"entry_task/pkg/logger"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
	"entry_task/pkg/page"
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

type ActivityTypeDao struct {
}

//查询公共sql
func (d ActivityTypeDao) querySql() *xorm.Session {
	return dao.SqlDB.NewSession().Table("activity_type_tab")
}

// FindList 查询集合
func (d ActivityTypeDao) FindList(query *req.ActivityTypeListQuery) ([]*response.ActivityTypeResponse, int64) {
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