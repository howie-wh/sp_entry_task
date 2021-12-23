package router

import (
	v1 "entry_task/admin/api/v1"
	"github.com/gin-gonic/gin"
)

//用户路由
func initUserRouter(router *gin.RouterGroup) {
	userApi := new(v1.UserApi)
	group := router.Group("/")
	{
		group.GET("/user_list", userApi.List)
	}
}
