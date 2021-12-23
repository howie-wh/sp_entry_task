package table

import (
	"reflect"
)

// Comment 用户表数据结构体
type Comment struct {
	CommentId   uint64    `xorm:"pk autoincr" json:"commentId"`     //评论ID
	UserId      uint64    `json:"userId"`      						//用户ID
	ActivityId  uint64    `json:"activityId"`      					//活动ID
	Content     string    `xorm:"varchar(256)" json:"content"`   	//评论内容
	DelFlag     string    `xorm:"char(1)" json:"delFlag"`         	//0正常1删除
	CreateTime  uint64    `xorm:"created" json:"createTime"`      	//创建时间
	UpdateTime  uint64    `json:"updateTime"`                     	//更新时间
}

func (Comment) TableName() string {
	return "comment_tab"
}

func (c Comment) IsEmpty() bool {
	return reflect.DeepEqual(c, Comment{})
}