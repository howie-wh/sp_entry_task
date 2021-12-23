package router

import (
	v1 "entry_task/activity/api/v1"
	"github.com/gin-gonic/gin"
)

//用户路由
func initActivityRouter(router *gin.RouterGroup) {
	activityApi := new(v1.ActivityApi)
	userRouter := router.Group("/")
	{
		//获取活动列表
		userRouter.GET("/activity_list", activityApi.List)
		//获取活动详情
		userRouter.GET("/activity_info", activityApi.GetInfo)
		//获取用户活动列表
		userRouter.GET("/user_activity_list", activityApi.UserActivityList)
	}
}
