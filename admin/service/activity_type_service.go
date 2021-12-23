package service

import (
	"entry_task/admin/dao"
	"entry_task/pkg/models/req"
	"entry_task/pkg/models/response"
	"entry_task/pkg/table"
)

// ActivityTypeService 业务逻辑
type ActivityTypeService struct {
	activityTypeDao dao.ActivityTypeDao
	activityDao dao.ActivityDao
}

// FindList 查询用户集合业务方法
func (s ActivityTypeService) FindList(query req.ActivityTypeListQuery) ([]*response.ActivityTypeResponse, int64) {
	return s.activityTypeDao.FindList(query)
}

// GetActivityTypeById 根据id查询用户数据
func (s ActivityTypeService) GetActivityTypeById(activityTypeId uint64) *response.ActivityTypeResponse {
	return s.activityTypeDao.GetActivityTypeById(activityTypeId)
}

// GetActivityTypeByTypeName 根据类型名称查询活动类型
func (s ActivityTypeService) GetActivityTypeByTypeName(typeName string) *table.ActivityType {
	activityType := table.ActivityType{}
	activityType.TypeName = typeName
	return s.activityTypeDao.GetActivityTypeByTypeName(activityType)
}

// Insert 添加业务逻辑
func (s ActivityTypeService) Insert(body req.ActivityTypeBody) bool {
	user := s.activityTypeDao.Insert(body)
	if user != nil {
		return true
	}
	return false
}

// Update 修改用户数据
func (s ActivityTypeService) Update(body req.ActivityTypeBody) int64 {
	return s.activityTypeDao.Update(body)
}

// Remove 根据id删除用户相关数据
// 活动类型表被活动表关联，所以需要删除关联的活动表
func (s ActivityTypeService) Remove(typeIds []uint64) int64 {
	// 删除对于关联的活动表
	if s.activityDao.RemoveByTypeIds(typeIds) == 0 {
		return 0
	}
	// 删除对于的活动类型表
	return s.activityTypeDao.Remove(typeIds)
}
