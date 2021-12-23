package table

import (
	"reflect"
)

// User 用户表数据结构体
type User struct {
	UserId      uint64    `xorm:"pk autoincr" json:"userId"`      //用户ID
	UserName    string    `xorm:"varchar(128)" json:"userName"`   //登录用户名
	Password    string    `xorm:"varchar(128)" json:"password"`   //密码
	NickName    string    `xorm:"varchar(128)" json:"nickName"`   //用户昵称
	Email       string    `xorm:"varchar(128)" json:"email"`      //邮箱
	Avatar      string    `xorm:"varchar(128)" json:"avatar"`     //头像路径
	DelFlag     string    `xorm:"char(1)" json:"delFlag"`         //0正常1删除
	CreateTime  uint64    `xorm:"created" json:"createTime"`      //创建时间
	UpdateTime  uint64    `json:"updateTime"`                     //更新时间
}

func (User) TableName() string {
	return "user_tab"
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(u, User{})
}
