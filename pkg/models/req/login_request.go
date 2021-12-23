package req

// LoginBody 登录参数
type LoginBody struct {
	UserName string `json:"userName" binding:"required"` //用户名
	Password string `json:"password" binding:"required"` //密码
}
