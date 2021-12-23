package v1

import (
	"entry_task/activity/service"
	"entry_task/pkg/models/req"
	"entry_task/pkg/page"
	"entry_task/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommentApi 活动类型操作api
type CommentApi struct {
	commentService service.CommentService
}

// List 查询活动类型列表
func (a CommentApi) List(c *gin.Context) {
	query := &req.CommentListQuery{}
	if c.BindQuery(&query) == nil {
		list, total := a.commentService.FindList(query)
		success := resp.Success(page.Page{
			Total: total,
			List:  list,
		}, "查询成功")
		c.JSON(200, success)
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}

// Publish 查询活动类型信息
func (a CommentApi) Publish(c *gin.Context) {
	body := &req.CommentBody{}
	if c.BindJSON(body) == nil {
		//添加评论
		UserId, _ := c.Get("UserId")
		if UserId == nil {
			c.JSON(http.StatusOK, resp.ErrorResp(http.StatusInternalServerError, "用户未登陆"))
			return
		}
		body.UserId = UserId.(uint64)
		if a.commentService.Insert(body) {
			c.JSON(http.StatusOK, resp.Success(nil))
		} else {
			c.JSON(http.StatusInternalServerError, resp.ErrorResp("保存失败"))
		}
	} else {
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("参数错误"))
	}
}