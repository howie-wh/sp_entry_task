package table

import (
	"reflect"
)

// SysUser 系统用户表数据结构体
type SysUser struct {
	UserId      uint64    `xorm:"pk autoincr" json:"userId"`     	//用户ID
	UserName    string    `xorm:"varchar(128)" json:"userName"`   	//用户名
	Password    string    `xorm:"varchar(128)" json:"password"`   	//用户秘密
	DelFlag     string    `xorm:"char(1)" json:"delFlag"`         	//0正常1删除
	CreateTime  uint64    `xorm:"created" json:"createTime"`      	//创建时间
	UpdateTime  uint64    `json:"updateTime"`                     	//更新时间
}

func (SysUser) TableName() string {
	return "sys_user_tab"
}

func (su SysUser) IsEmpty() bool {
	return reflect.DeepEqual(su, SysUser{})
}
