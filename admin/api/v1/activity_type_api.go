package v1

import (
	"entry_task/admin/service"
	"entry_task/pkg/models/req"
	"entry_task/pkg/page"
	"entry_task/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ActivityTypeApi 活动类型操作api
type ActivityTypeApi struct {
	activityTypeService service.ActivityTypeService
}

// List 查询活动类型列表
func (a ActivityTypeApi) List(c *gin.Context) {
	query := req.ActivityTypeListQuery{}
	if c.BindQuery(&query) == nil {
		list, total := a.activityTypeService.FindList(query)
		success := resp.Success(page.Page{
			Total: total,
			List:  list,
		}, "查询成功")
		c.JSON(200, success)
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}

// GetInfo 查询活动类型信息
func (a ActivityTypeApi) GetInfo(c *gin.Context) {
	query := req.ActivityTypeQuery{}
	if c.BindQuery(&query) == nil {
		activityType := a.activityTypeService.GetActivityTypeById(query.TypeId)
		if activityType != nil {
			success := resp.Success(activityType, "查询成功")
			c.JSON(200, success)
		} else {
			c.JSON(200, resp.ErrorResp())
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}

// Add 新增活动类型
func (a ActivityTypeApi) Add(c *gin.Context) {
	body := req.ActivityTypeBody{}
	if c.BindJSON(&body) == nil {
		//根据类型名称查询活动类型
		activityType := a.activityTypeService.GetActivityTypeByTypeName(body.TypeName)
		if activityType != nil {
			if activityType.DelFlag == "1" {
				body.TypeId = activityType.TypeId
				body.DelFlag = "0"
				if a.activityTypeService.Update(body) != 0 {
					c.JSON(http.StatusOK, resp.Success(nil))
				} else {
					c.JSON(http.StatusInternalServerError, resp.ErrorResp("保存失败"))
				}
			} else {
				c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "失败，活动类型已存在"))
			}
			return
		}
		//添加活动类型
		if a.activityTypeService.Insert(body) {
			c.JSON(http.StatusOK, resp.Success(nil))
		} else {
			c.JSON(http.StatusInternalServerError, resp.ErrorResp("保存失败"))
		}
	} else {
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("参数错误"))
	}
}

// Update 修改活动类型
func (a ActivityTypeApi) Update(c *gin.Context) {
	body := req.ActivityTypeBody{}
	if c.BindJSON(&body) == nil {
		//根据类型名称查询活动类型
		activityType := a.activityTypeService.GetActivityTypeByTypeName(body.TypeName)
		if activityType != nil {
			c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "失败，活动类型已存在"))
			return
		}
		//进行用户修改操作
		if a.activityTypeService.Update(body) > 0 {
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

// Remove 删除活动类型
func (a ActivityTypeApi) Remove(c *gin.Context) {
	body := req.ActivityTypeDeleteBody{}
	if c.BindJSON(&body) == nil {
		if a.activityTypeService.Remove(body.TypeIds) > 0 {
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