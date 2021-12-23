package v1

import (
	"entry_task/activity/service"
	"entry_task/pkg/models/req"
	"entry_task/pkg/page"
	"entry_task/pkg/resp"
	"github.com/gin-gonic/gin"
)

// ActivityApi 操作api
type ActivityApi struct {
	activityService service.ActivityService
}

// List 查询活动列表
func (a ActivityApi) List(c *gin.Context) {
	query := &req.ActivityListQuery{}
	if c.BindQuery(query) == nil {
		list, total := a.activityService.FindList(query)
		success := resp.Success(page.Page{
			Total: total,
			List:  list,
		}, "查询成功")
		c.JSON(200, success)
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}

// GetInfo 查询活动详情信息
func (a ActivityApi) GetInfo(c *gin.Context) {
	query := &req.ActivityQuery{}
	if c.BindQuery(query) == nil {
		// 获取登陆用户id
		UserId, _ := c.Get("UserId")
		if UserId != nil {
			query.UserId = UserId.(uint64)
		}
		activity := a.activityService.GetActivityById(query)
		if activity != nil {
			success := resp.Success(activity, "查询成功")
			c.JSON(200, success)
		} else {
			c.JSON(200, resp.ErrorResp())
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}

// UserActivityList 查询用户活动列表
func (a ActivityApi) UserActivityList(c *gin.Context) {
	query := &req.ActivityListQuery{}
	if c.BindQuery(query) == nil {
		// 获取登陆用户id
		UserId, _ := c.Get("UserId")
		if UserId == nil {
			c.JSON(200, resp.ErrorResp(500, "用户未登陆"))
			return
		}
		query.UserId = UserId.(uint64)
		list, total := a.activityService.FindUserActivityList(query)
		success := resp.Success(page.Page{
			Total: total,
			List:  list,
		}, "查询成功")
		c.JSON(200, success)
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}