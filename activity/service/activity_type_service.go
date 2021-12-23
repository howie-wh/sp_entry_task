package service

import (
	"entry_task/activity/dao"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
)

// ActivityTypeService 业务逻辑
type ActivityTypeService struct {
	activityTypeDao dao.ActivityTypeDao
}

// FindList 查询用户集合业务方法
func (s ActivityTypeService) FindList(query *req.ActivityTypeListQuery) ([]*response.ActivityTypeResponse, int64) {
	return s.activityTypeDao.FindList(query)
}