package router

import (
	v1 "entry_task/activity/api/v1"
	"github.com/gin-gonic/gin"
)

//用户路由
func initCommentRouter(router *gin.RouterGroup) {
	commentApi := new(v1.CommentApi)
	userRouter := router.Group("/")
	{
		//获取评论列表
		userRouter.GET("/comment_list", commentApi.List)
		//发布评论
		userRouter.POST("/comment_publish", commentApi.Publish)
	}
}
