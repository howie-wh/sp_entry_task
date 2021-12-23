package table

import (
	"reflect"
)

// ActivityJoin 用户表数据结构体
type ActivityJoin struct {
	JoinId      uint64    `xorm:"pk autoincr" json:"joinId"`      	//报名ID
	UserId      uint64    `json:"userId"`      						//用户ID
	ActivityId  uint64    `json:"activityId"`      					//活动ID
	DelFlag     string    `xorm:"char(1)" json:"delFlag"`         	//0正常1删除
	CreateTime  uint64    `xorm:"created" json:"createTime"`      	//创建时间
	UpdateTime  uint64    `json:"updateTime"`                     	//更新时间
}

func (ActivityJoin) TableName() string {
	return "activity_join_tab"
}

func (aj ActivityJoin) IsEmpty() bool {
	return reflect.DeepEqual(aj, ActivityJoin{})
}
