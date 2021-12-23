package router

import (
	v1 "entry_task/activity/api/v1"
	"github.com/gin-gonic/gin"
)

//用户路由
func initUserRouter(router *gin.RouterGroup) {
	userApi := new(v1.UserApi)
	userRouter := router.Group("/")
	{
		//获取用户列表
		userRouter.GET("/user_list", userApi.List)
		//获取用户详情
		userRouter.GET("/user_info", userApi.GetInfo)
		//用户注册
		userRouter.POST("/user_register", userApi.Register)
	}
}
