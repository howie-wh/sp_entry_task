package router

import (
	v1 "entry_task/activity/api/v1"
	"github.com/gin-gonic/gin"
)

//用户路由
func initActivityTypeRouter(router *gin.RouterGroup) {
	activityTypeApi := new(v1.ActivityTypeApi)
	group := router.Group("/")
	{
		//获取活动类型列表
		group.GET("/activity_type_list", activityTypeApi.List)
	}
}
