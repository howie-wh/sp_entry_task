package req

import (
	"entry_task/pkg/base"
)

// ActivityTypeListQuery 用户get请求数据参数
type ActivityTypeListQuery struct {
	base.GlobalPageQuery
	TypeName   string  	`form:"typeName"`      		//活动类型名称
}

// ActivityTypeQuery 用户get请求数据参数
type ActivityTypeQuery struct {
	TypeId   	uint64  	`form:"typeId" binding:"required"`      //活动类型id
}

// ActivityTypeBody 用户接收POST 或者 PUT请求参数
type ActivityTypeBody struct {
	TypeId  	uint64     	`xorm:"pk autoincr" json:"typeId"` 				//活动ID
	TypeName    string    	`json:"typeName" binding:"required,max=64"`    	//活动标题
	DelFlag     string    	`json:"delFlag"`                   				//0正常1删除
	CreateTime  uint64 	    `xorm:"created" json:"createTime"` 				//创建时间
	UpdateTime  uint64 	    `json:"updateTime"`                				//更新时间
}

// ActivityTypeDeleteBody 用户接收delete请求参数
type ActivityTypeDeleteBody struct {
	TypeIds  	[]uint64     `json:"typeIds" binding:"required"` 	//活动类型ID列表
}
