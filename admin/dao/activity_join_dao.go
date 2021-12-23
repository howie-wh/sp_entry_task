package dao

import (
	"entry_task/pkg/dao"
	"entry_task/pkg/logger"
	"entry_task/pkg/table"
	"github.com/sirupsen/logrus"
	"time"
)

type ActivityJoinDao struct {
}

// RemoveByActivityIds 根据id删除数据
func (d ActivityJoinDao) RemoveByActivityIds(activityIds []uint64) int64 {
	tab := &table.ActivityJoin{
		DelFlag: "1",
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession().Table("activity_join_tab")
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
