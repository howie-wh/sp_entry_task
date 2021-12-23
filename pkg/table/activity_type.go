package table

import (
	"reflect"
)

// ActivityType 活动类型数据结构体
type ActivityType struct {
	TypeId  	uint64    `xorm:"pk autoincr" json:"typeId"`     //活动类型ID
	TypeName    string    `xorm:"varchar(64)" json:"typeName"`   //活动类型名称
	DelFlag     string    `xorm:"char(1)" json:"delFlag"`        //0正常1删除
	CreateTime  uint64    `xorm:"created" json:"createTime"`     //创建时间
	UpdateTime  uint64    `json:"updateTime"`                    //更新时间
}

func (ActivityType) TableName() string {
	return "activity_type_tab"
}

func (at ActivityType) IsEmpty() bool {
	return reflect.DeepEqual(at, ActivityType{})
}