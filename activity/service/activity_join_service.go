package service

import (
	"entry_task/activity/dao"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
)

// ActivityJoinService 业务逻辑
type ActivityJoinService struct {
	activityJoinDao dao.ActivityJoinDao
}

// Insert 添加业务逻辑
func (s ActivityJoinService) Insert(body *req.ActivityJoinBody) bool {
	user := s.activityJoinDao.Insert(body)
	if user != nil {
		return true
	}
	return false
}

// Update 修改数据
func (s ActivityJoinService) Update(body *req.ActivityJoinBody) int64 {
	return s.activityJoinDao.Update(body)
}

// Remove 根据用户id删除用户相关数据
func (s ActivityJoinService) Remove(activityId, userId uint64) int64 {
	return s.activityJoinDao.Remove(activityId, userId)
}

// GetActivityJoinById 根据id查询用户数据
func (s ActivityJoinService) GetActivityJoinById(activityId, userId uint64) *response.ActivityJoinResponse {
	return s.activityJoinDao.GetActivityJoinById(activityId, userId)
}