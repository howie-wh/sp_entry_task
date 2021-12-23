package v1

import (
	"entry_task/activity/service"
	"entry_task/pkg/models/req"
	"entry_task/pkg/page"
	"entry_task/pkg/resp"
	"github.com/gin-gonic/gin"
)

// ActivityTypeApi 活动类型操作api
type ActivityTypeApi struct {
	activityTypeService service.ActivityTypeService
}

// List 查询活动类型列表
func (a ActivityTypeApi) List(c *gin.Context) {
	query := &req.ActivityTypeListQuery{}
	if c.BindQuery(&query) == nil {
		list, i := a.activityTypeService.FindList(query)
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