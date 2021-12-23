package v1

import (
	"entry_task/activity/service"
	"entry_task/pkg/models/req"
	"entry_task/pkg/page"
	"entry_task/pkg/resp"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserApi 用户操作api
type UserApi struct {
	userService service.UserService
}

// List 查询用户列表
func (a UserApi) List(c *gin.Context) {
	query := &req.UserListQuery{}
	if c.BindQuery(query) == nil {
		list, total := a.userService.FindList(query)
		success := resp.Success(page.Page{
			Total: total,
			List:  list,
		}, "查询成功")
		c.JSON(200, success)
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}

// GetInfo 查询用户信息
func (a UserApi) GetInfo(c *gin.Context) {
	//获取用户id
	UserId, _ := c.Get("UserId")
	if UserId == nil {
		c.JSON(http.StatusOK, resp.ErrorResp("用户未登陆"))
		return
	}
	query := &req.UserQuery{}
	if c.BindQuery(query) == nil {
		user := a.userService.GetUserById(UserId.(uint64))
		if user != nil {
			success := resp.Success(user, "查询成功")
			c.JSON(200, success)
		} else {
			c.JSON(200, resp.ErrorResp())
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}

// Register 用户注册
func (a UserApi) Register(c *gin.Context) {
	body := &req.UserBody{}
	if c.BindJSON(body) == nil {
		//根据用户名查询用户
		user := a.userService.GetUserByUserName(body.UserName)
		if user != nil {
			c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "失败，登录账号已存在"))
			return
		}
		//进行密码加密
		body.Password = gotool.BcryptUtils.Generate(body.Password)
		//添加用户
		if a.userService.Insert(body) {
			c.JSON(http.StatusOK, resp.Success(nil))
		} else {
			c.JSON(http.StatusInternalServerError, resp.ErrorResp("保存失败"))
		}
	} else {
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("参数错误"))
	}
}