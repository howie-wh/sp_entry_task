package jwt

import (
	"entry_task/pkg/constant"
	"entry_task/pkg/dao"
	"entry_task/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserUtils struct {
}

// GetUserInfo 通过jwt获取当前登录用户
func GetUserInfo(c *gin.Context) *CustomUserInfo {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		return nil
	}
	j := NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		return nil
	}
	info := claims.UserInfo
	return &info
}

// SaveRedisToken 将token存入到redis
func SaveRedisToken(key string, s string) {
	dao.RedisDB.SETEX(key, constant.RedisConstant{}.GetRedisTokenExpires(), s)
}

// RemoveKeyFromRedis 根据key删除
func RemoveKeyFromRedis(key string) int {
	del, err := dao.RedisDB.DEL(key)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
	}
	return del
}