package router

import (
	v1 "entry_task/admin/api/v1"
	"github.com/gin-gonic/gin"
)

//用户路由
func initActivityTypeRouter(router *gin.RouterGroup) {
	activityTypeApi := new(v1.ActivityTypeApi)
	group := router.Group("/")
	{
		//获取活动类型列表
		group.GET("/activity_type_list", activityTypeApi.List)
		//获取活动类型详情
		group.GET("/activity_type", activityTypeApi.GetInfo)
		//新增活动类型
		group.POST("/activity_type", activityTypeApi.Add)
		//修改活动类型
		group.PUT("/activity_type", activityTypeApi.Update)
		//删除活动类型
		group.DELETE("/activity_type", activityTypeApi.Remove)
	}
}
