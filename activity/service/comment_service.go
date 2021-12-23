package service

import (
	"entry_task/activity/dao"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
)

// CommentService 业务逻辑
type CommentService struct {
	commentDao dao.CommentDao
}

// FindList 查询用户集合业务方法
func (s CommentService) FindList(query *req.CommentListQuery) ([]*response.CommentResponse, int64) {
	return s.commentDao.FindList(query)
}

// Insert 发布评论
func (s CommentService) Insert(body *req.CommentBody) bool {
	c := s.commentDao.Insert(body)
	if c != nil {
		return true
	}
	return false
}