package req

import (
	"entry_task/pkg/base"
)

// UserListQuery 用户get请求数据参数
type UserListQuery struct {
	base.GlobalPageQuery
}

type UserQuery struct {
}

// UserBody 用户接收POST 或者 PUT请求参数
type UserBody struct {
	UserId      uint64    `xorm:"pk autoincr" json:"userId"` 						//用户ID
	UserName    string    `json:"userName" binding:"required,min=2,max=128"`      	//登录用户名
	Password    string    `json:"password" binding:"required,min=6,max=128"`      	//密码
	NickName    string    `json:"nickName" binding:"max=128"`                  		//用户昵称
	Email       string    `json:"email" binding:"email,max=128"`            		//邮箱
	Avatar      string    `json:"avatar" binding:"url,max=128"`             		//头像路径
	DelFlag     string    `json:"delFlag"`                   						//0正常1删除
	CreateTime  uint64 	  `xorm:"created" json:"createTime"` 						//创建时间
	UpdateTime  uint64 	  `json:"updateTime"`                						//更新时间
}
