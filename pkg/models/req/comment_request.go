package req

import (
	"entry_task/pkg/base"
)

// CommentListQuery 用户get请求数据参数
type CommentListQuery struct {
	base.GlobalPageQuery
	UserId  	uint64  	`form:"userId"`     	//用户id
	ActivityId 	uint64    	`form:"activityId"`     //活动id
}

// CommentQuery 用户get请求数据参数
type CommentQuery struct {
	CommentId  uint64  	`form:"commentId" binding:"required"`     //评论id
}

// CommentBody 用户发布评论
type CommentBody struct {
	CommentId   uint64     	`xorm:"pk autoincr" json:"commentId"` 		//评论ID
	ActivityId  uint64     	`json:"activityId" binding:"required"` 		//活动ID
	UserId      uint64     	`json:"userId" binding:"required"`      	//用户ID
	Content     string    	`json:"content" binding:"required,max=255"` //评论内容
	DelFlag     string    	`json:"delFlag"`                   			//0正常1删除
	CreateTime  uint64 	    `xorm:"created" json:"createTime"` 			//创建时间
	UpdateTime  uint64 	    `json:"updateTime"`                			//更新时间
}
