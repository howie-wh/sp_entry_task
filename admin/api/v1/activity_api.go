package v1

import (
	"entry_task/admin/service"
	"entry_task/pkg/models/req"
	"entry_task/pkg/page"
	"entry_task/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ActivityApi 操作api
type ActivityApi struct {
	activityService service.ActivityService
}

// List 查询用户列表
func (a ActivityApi) List(c *gin.Context) {
	query := req.ActivityListQuery{}
	if c.BindQuery(&query) == nil {
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

// GetInfo 查询活动信息
func (a ActivityApi) GetInfo(c *gin.Context) {
	query := req.ActivityQuery{}
	if c.BindQuery(&query) == nil {
		activity := a.activityService.GetActivityById(query.ActivityId)
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

// Add 新增活动
func (a ActivityApi) Add(c *gin.Context) {
	body := req.ActivityBody{}
	if c.BindJSON(&body) == nil {
		//根据标题查询活动
		activity := a.activityService.GetActivityByTitle(body.Title)
		if activity != nil {
			if activity.DelFlag == "1" {
				body.ActivityId = activity.ActivityId
				body.DelFlag = "0"
				if a.activityService.Update(body) != 0 {
					c.JSON(http.StatusOK, resp.Success(nil))
				} else {
					c.JSON(http.StatusInternalServerError, resp.ErrorResp("保存失败"))
				}
			} else {
				c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "失败，活动标题已存在"))
			}
			return
		}
		//添加活动
		if a.activityService.Insert(body) {
			c.JSON(http.StatusOK, resp.Success(nil))
		} else {
			c.JSON(http.StatusInternalServerError, resp.ErrorResp("保存失败"))
		}
	} else {
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("参数错误"))
	}
}

// Update 修改用户
func (a ActivityApi) Update(c *gin.Context) {
	body := req.ActivityBody{}
	if c.BindJSON(&body) == nil {
		//进行用户修改操作
		if a.activityService.Update(body) > 0 {
			resp.OK(c)
			return
		} else {
			resp.Error(c)
			return
		}
	} else {
		resp.ParamError(c)
		return
	}
}

// Remove 删除活动
func (a ActivityApi) Remove(c *gin.Context) {
	body := req.ActivityDeleteBody{}
	if c.BindJSON(&body) == nil {
		if a.activityService.Remove(body.ActivityIds) > 0 {
			resp.OK(c)
			return
		} else {
			resp.Error(c)
			return
		}
	} else {
		resp.ParamError(c)
		return
	}
}