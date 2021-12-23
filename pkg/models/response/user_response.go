package response

// UserResponse 用户实体返回结构体
type UserResponse struct {
	UserId      uint64         `json:"userId"`           //用户ID
	UserName    string         `json:"userName"`         //登录用户名
	NickName    string         `json:"nickName"`         //用户昵称
	Email       string         `json:"email"`            //邮箱
	Avatar      string         `json:"avatar"`           //头像路径
}

// SysUserResponse 用户实体返回结构体
type SysUserResponse struct {
	UserId      uint64         `json:"userId"`           //用户ID
	UserName    string         `json:"userName"`         //登录用户名
}

