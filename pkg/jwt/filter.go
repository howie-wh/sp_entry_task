package jwt

import (
	"entry_task/pkg/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// X 全局
var (
	WhiteList []string
	BlackList []string
)

//是否在放行范围内
func doSquare(c *gin.Context) bool {
	for i := 0; i < len(WhiteList); i++ {
		replace := strings.Contains(c.Request.RequestURI, WhiteList[i])
		if replace {
			return true
		}
	}
	return false
}

// AddWhiteList 放行的请求
func AddWhiteList(uri string)  {
	WhiteList = append(WhiteList, uri)
}

// AddBlackList 放行的请求
func AddBlackList(uri string)  {
	BlackList = append(BlackList, uri)
}

// JWTFilter 中间件，检查token
func JWTFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		//调用过滤去将放行的请求先放行
		if doSquare(c) {
			return
		}
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": 401,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 401,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		get, err := dao.RedisDB.GET(claims.UserInfo.UserName)
		if err == nil {
			if !(get == token) {
				c.JSON(http.StatusOK, gin.H{
					"status": 401,
					"msg":    "您的账号已在其他终端登录，请重新登录",
				})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 401,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		// c.Set("claims", claims)
	}
}

// JWTCheckUser 判断是否是已经登陆
// 如果已经登陆，将用户名称设置Context的key中
func JWTCheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := GetUserInfo(c)
		if userInfo == nil {
			return
		}
		get, err := dao.RedisDB.GET(userInfo.UserName)
		if err != nil {
			return
		}
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			return
		} else if get != token {
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("UserId", userInfo.UserId)
		c.Set("UserName", userInfo.UserName)
	}
}