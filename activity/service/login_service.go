package service

import (
	"entry_task/pkg/jwt"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"strings"
)

type LoginService struct {
	userService UserService
}

// Login 用户登录业务处理
func (s LoginService) Login(name string, password string) (bool, string) {
	user := s.userService.GetUserByUserName(name)
	if user == nil {
		return false, "用户不存在"
	}
	if !gotool.BcryptUtils.CompareHash(user.Password, password) {
		return false, "密码错误"
	}
	//生成token
	userInfo := s.userService.GetUserById(user.UserId)
	customUserInfo := &jwt.CustomUserInfo{
		UserId: userInfo.UserId,
		UserName: userInfo.UserName,
	}
	token, err := jwt.NewJWT().CreateUserToken(customUserInfo)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return false, ""
	}

	//将token存入到redis中
	jwt.SaveRedisToken(name, token)
	return true, token
}

// LoginUser 获取当前登录用户
func (s LoginService) LoginUser(c *gin.Context) *jwt.CustomUserInfo {
	token := c.Request.Header.Get("Authorization")
	str := strings.Split(token, " ")
	j := jwt.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(str[1])
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
	}
	info := claims.UserInfo
	return &info
}

// Logout 获取当前登录用户
func (s LoginService) Logout(c *gin.Context) error {
	name := jwt.GetUserInfo(c).UserName
	jwt.RemoveKeyFromRedis(name)
	return nil
}