package router

import (
	v1 "entry_task/admin/api/v1"
	"github.com/gin-gonic/gin"
)

//用户路由
func initActivityRouter(router *gin.RouterGroup) {
	activityApi := new(v1.ActivityApi)
	group := router.Group("/")
	{
		//获取活动列表
		group.GET("/activity_list", activityApi.List)
		//获取活动详情
		group.GET("/activity", activityApi.GetInfo)
		//新增活动
		group.POST("/activity", activityApi.Add)
		//修改活动
		group.PUT("/activity", activityApi.Update)
		//删除活动
		group.DELETE("/activity", activityApi.Remove)
	}
}
