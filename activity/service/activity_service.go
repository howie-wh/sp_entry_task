package service

import (
	//"pkg/table"
	"entry_task/activity/dao"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
)

// ActivityService 用户操作业务逻辑
type ActivityService struct {
	activityDao dao.ActivityDao
	activityJoinDao dao.ActivityJoinDao
}

// FindList 查询活动列表
func (s ActivityService) FindList(query *req.ActivityListQuery) ([]*response.ActivityResponse, int64) {
	return s.activityDao.FindList(query)
}

// GetActivityById 根据id查询活动数据
func (s ActivityService) GetActivityById(query *req.ActivityQuery) *response.ActivityResponse {
	return s.activityDao.GetActivityById(query)
}

// FindUserActivityList 查询用户活动列表
func (s ActivityService) FindUserActivityList(query *req.ActivityListQuery) ([]*response.ActivityResponse, int64) {
	list, total := s.activityDao.FindList(query)
	if list == nil {
		return nil, 0
	}
	for _, resp := range list {
		r := s.activityJoinDao.GetActivityJoinById(resp.ActivityId, query.UserId)
		if r != nil {
			resp.IsJoin = true
		}
	}
	return list, total
	//return s.activityDao.FindUserActivityList(query)
}