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
	"time"
)

type CommentDao struct {
}

//查询公共sql
func (d CommentDao) querySql() *xorm.Session {
	return dao.SqlDB.NewSession().Table("comment_tab")
}

// FindList 查询集合
func (d CommentDao) FindList(query *req.CommentListQuery) ([]*response.CommentResponse, int64) {
	resp := make([]*response.CommentResponse, 0)
	sql := d.querySql()
	if query.UserId > 0 {
		sql.And("user_id = ?", query.UserId)
	}
	if query.ActivityId > 0 {
		sql.And("activity_id = ?", query.ActivityId)
	}
	sql.And("del_flag = 0")

	total, _ := page.GetTotal(sql.Clone())
	err := sql.Desc("create_time").Limit(query.Offset, query.Start).Find(&resp)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil, 0
	}
	return resp, total
}

// GetCommentById 根据id查询数据
func (d CommentDao) GetCommentById(CommentId uint64) *response.CommentResponse {
	var resp response.CommentResponse
	get, err := d.querySql().Where("comment_id = ?", CommentId).Get(&resp)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return nil
	}
	if !get {
		return nil
	}
	return &resp
}

// Insert 添加用户活动评论
func (d CommentDao) Insert(body *req.CommentBody) *table.Comment {
	tab := &table.Comment{
		UserId: body.UserId,
		ActivityId: body.ActivityId,
		Content: body.Content,
		DelFlag: "0",
		UpdateTime: uint64(time.Now().Unix()),
	}
	session := dao.SqlDB.NewSession()
	session.Begin()
	_, err := session.Table("comment_tab").Insert(tab)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		session.Rollback()
		return nil
	}
	session.Commit()
	return tab
}
