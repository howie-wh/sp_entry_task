package jwt

import (
	"entry_task/pkg/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"

)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("授权已过期")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("令牌非法")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "0df9b8db-6f7c-d713-eeab-ecb317696042"
)

// CustomUserInfo redis中保存的用户信息，兼容普通用户和系统用户
type CustomUserInfo struct {
	UserId      uint64         `json:"userId"`           //用户ID
	UserName    string         `json:"userName"`         //登录用户名
	UserMode    string         `json:"userMode"`         //用户模式,1-admin, 0-normal
}

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserInfo CustomUserInfo `json:"userInfo"`
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// GetSignKey 获取signKey
func GetSignKey() string {
	return SignKey
}

// SetSignKey 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateUserToken 生成含有用户信息的token
func (j *JWT) CreateUserToken(u *CustomUserInfo) (string, error) {
	jwtConfig := config.GetJwtConfig()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserInfo: *u,
		StandardClaims: jwt.StandardClaims{
			//设置一小时时效
			ExpiresAt: time.Now().Add(jwtConfig.TimeOut * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    jwtConfig.Issuer,
		},
	})
	return claims.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}