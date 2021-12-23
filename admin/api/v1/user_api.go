package v1

import (
	"entry_task/admin/service"
	"entry_task/pkg/models/req"
	"entry_task/pkg/page"
	"entry_task/pkg/resp"
	"github.com/gin-gonic/gin"
)

// UserApi 用户操作api
type UserApi struct {
	userService service.UserService
}

// List 查询用户列表
func (a UserApi) List(c *gin.Context) {
	query := req.UserListQuery{}
	if c.BindQuery(&query) == nil {
		list, i := a.userService.FindList(query)
		success := resp.Success(page.Page{
			//Size:  query.Offset,
			Total: i,
			List:  list,
		}, "查询成功")
		c.JSON(200, success)
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}