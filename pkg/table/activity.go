package table

import (
	"reflect"
)

// Activity 活动表数据结构体
type Activity struct {
	ActivityId  uint64    `xorm:"pk autoincr" json:"activityId"`    //活动ID
	TypeId      uint64    `json:"typeId"`                         	//活动类型ID
	Title    	string    `xorm:"varchar(128)" json:"title"`   		//活动标题
	Location    string    `xorm:"varchar(128)" json:"location"`   	//活动地址
	Content     string    `xorm:"varchar(1024)" json:"content"`     //活动内容
	StartTime  	uint64    `json:"startTime"`      					//开始时间
	EndTime  	uint64    `json:"endTime"`                     		//结束时间
	DelFlag     string    `xorm:"char(1)" json:"delFlag"`         	//0正常1删除
	CreateTime  uint64    `xorm:"created" json:"createTime"`      	//创建时间
	UpdateTime  uint64    `json:"updateTime"`                     	//更新时间
}

func (Activity) TableName() string {
	return "activity_tab"
}

func (a Activity) IsEmpty() bool {
	return reflect.DeepEqual(a, Activity{})
}
