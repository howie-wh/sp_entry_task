package service

import (
	"entry_task/admin/dao"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
	"entry_task/pkg/table"
)

// ActivityService 用户操作业务逻辑
type ActivityService struct {
	activityDao dao.ActivityDao
	activityJoinDao dao.ActivityJoinDao
}

// FindList 查询用户集合业务方法
func (s ActivityService) FindList(query req.ActivityListQuery) ([]*response.ActivityResponse, int64) {
	return s.activityDao.FindList(query)
}

// GetActivityById 根据id查询用户数据
func (s ActivityService) GetActivityById(activityId uint64) *response.ActivityResponse {
	return s.activityDao.GetActivityById(activityId)
}

// GetActivityByTitle 根据标题查询活动
func (s ActivityService) GetActivityByTitle(title string) *table.Activity {
	activity := table.Activity{}
	activity.Title = title
	return s.activityDao.GetActivityByTitle(activity)
}

// Insert 添加业务逻辑
func (s ActivityService) Insert(body req.ActivityBody) bool {
	user := s.activityDao.Insert(body)
	if user != nil {
		return true
	}
	return false
}

// Update 修改数据
func (s ActivityService) Update(body req.ActivityBody) int64 {
	return s.activityDao.Update(body)
}

// Remove 根据用户id删除用户相关数据
// 活动表被报名表依赖，所以需要删除对于报名表中关联的活动
func (s ActivityService) Remove(activityIds []uint64) int64 {
	// 删除对于关联的活动报名表
	if s.activityJoinDao.RemoveByActivityIds(activityIds) == 0 {
		return 0
	}
	// 删除对于的活动表
	return s.activityDao.Remove(activityIds)
}
