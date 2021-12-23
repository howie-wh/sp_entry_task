package req

import (
	"entry_task/pkg/base"
)

// ActivityListQuery 用户get请求数据参数
type ActivityListQuery struct {
	base.GlobalPageQuery
	StartTime   uint64  	`form:"startTime"`      //开始时间
	EndTime     uint64 		`form:"endTime"`    	//结束时间
	TypeId      uint64  	`form:"typeId"`      	//活动类型id
	UserId      uint64      `form:"userId"`         //用户id
}

// ActivityQuery 用户get请求数据参数
type ActivityQuery struct {
	ActivityId  uint64  	`form:"activityId" binding:"required"`     //活动id
	UserId      uint64      `form:"userId"`         //用户id
}

// ActivityBody 用户接收POST 或者 PUT请求参数
type ActivityBody struct {
	ActivityId  uint64     	`xorm:"pk autoincr" json:"activityId"` 		//活动ID
	TypeId      uint64     	`json:"typeId" binding:"required"`          //活动类型ID
	Title    	string    	`json:"title" binding:"required,max=128"`           //活动标题
	Location    string    	`json:"location" binding:"required,max=128"`        //活动地点
	Content     string    	`json:"content" binding:"required,max=1024"`         //活动内容
	StartTime  	uint64 	    `json:"startTime" binding:"required"` 		//开始时间
	EndTime  	uint64 	    `json:"endTime" binding:"required"`         //结束时间
	DelFlag     string    	`json:"delFlag"`                   		//0正常1删除
	CreateTime  uint64 	    `xorm:"created" json:"createTime"` 		//创建时间
	UpdateTime  uint64 	    `json:"updateTime"`                		//更新时间
}

// ActivityDeleteBody 用户接收delete请求参数
type ActivityDeleteBody struct {
	ActivityIds  []uint64     `json:"activityIds" binding:"required"` 		//活动ID列表
}
