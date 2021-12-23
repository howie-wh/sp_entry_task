package router

import (
	"entry_task/pkg/jwt"
	"entry_task/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()
	//router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Recover)
	router.Use(jwt.JWTFilter())
	router.Use(jwt.JWTCheckUser())
	//v1版本api
	v1Router := router.Group("/api/entry_task/v1/activity" +
		"")
	{
		//登录接口
		initLoginRouter(v1Router)
		//用户接口
		initUserRouter(v1Router)
		//活动接口
		initActivityRouter(v1Router)
		//活动类型接口
		initActivityTypeRouter(v1Router)
		//活动报名接口
		initActivityJoinRouter(v1Router)
		//评论接口
		initCommentRouter(v1Router)
	}
	return router
}
