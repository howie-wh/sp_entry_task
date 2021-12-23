package v1

import (
	"entry_task/activity/service"
	"entry_task/pkg/models/req"
	"entry_task/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ActivityJoinApi 活动报名操作api
type ActivityJoinApi struct {
	activityJoinService service.ActivityJoinService
}

// Join 新增活动类型
func (a ActivityJoinApi) Join(c *gin.Context) {
	body := &req.ActivityJoinBody{}
	if c.BindJSON(body) == nil {
		// 获取登陆用户id
		UserId, _ := c.Get("UserId")
		if UserId == nil {
			c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "用户未登陆"))
			return
		}
		body.UserId = UserId.(uint64)
		//根据类型名称查询活动类型
		activityJoin := a.activityJoinService.GetActivityJoinById(body.ActivityId, body.UserId)
		if activityJoin != nil {
			if activityJoin.DelFlag == "1" {
				body.JoinId = activityJoin.JoinId
				body.DelFlag = "0"
				if a.activityJoinService.Update(body) != 0 {
					c.JSON(http.StatusOK, resp.Success(nil))
				} else {
					c.JSON(http.StatusInternalServerError, resp.ErrorResp("保存失败"))
				}
			} else {
				c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "失败，活动已经报名"))
			}
			return
		}
		//添加活动类型
		if a.activityJoinService.Insert(body) {
			c.JSON(http.StatusOK, resp.Success(nil))
		} else {
			c.JSON(http.StatusInternalServerError, resp.ErrorResp("保存失败"))
		}
	} else {
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("参数错误"))
	}
}


// Quit 删除活动类型
func (a ActivityJoinApi) Quit(c *gin.Context) {
	body := &req.ActivityJoinBody{}
	if c.BindJSON(body) == nil {
		// 获取登陆用户id
		UserId, _ := c.Get("UserId")
		if UserId == nil {
			c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "用户未登陆"))
			return
		}
		if a.activityJoinService.Remove(body.ActivityId, UserId.(uint64)) > 0 {
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