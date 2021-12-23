package router

import (
	v1 "entry_task/activity/api/v1"
	"github.com/gin-gonic/gin"
)

//用户路由
func initActivityJoinRouter(router *gin.RouterGroup) {
	activityJoinApi := new(v1.ActivityJoinApi)
	userRouter := router.Group("/")
	{
		//参加活动
		userRouter.POST("/activity_join", activityJoinApi.Join)
		//退出活动
		userRouter.POST("/activity_quit", activityJoinApi.Quit)
	}
}
