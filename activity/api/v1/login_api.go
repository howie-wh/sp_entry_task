package v1

import (
	"entry_task/activity/service"
	"entry_task/pkg/models/req"
	"entry_task/pkg/resp"
	"github.com/gin-gonic/gin"
)

type LoginApi struct {
	loginService service.LoginService
}

// Login 登录
func (a LoginApi) Login(c *gin.Context) {
	loginBody := &req.LoginBody{}
	if c.BindJSON(loginBody) == nil {
		m := make(map[string]string)
		login, s := a.loginService.Login(loginBody.UserName, loginBody.Password)
		if login {
			m["token"] = s
			c.JSON(200, resp.Success(m))
		} else {
			c.JSON(200, resp.ErrorResp(s))
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数绑定错误"))
	}
}

// Logout 退出登录
func (a LoginApi) Logout(c *gin.Context) {
	err := a.loginService.Logout(c)
	if err != nil {
		c.JSON(200, resp.ErrorResp())
	}
	resp.OK(c)
}
