package model

import "github.com/dgrijalva/jwt-go"

// 模型层，放实例结构体
type User struct {
	Username string `form:"Username" json:"Username" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
}
type Myclaims struct {
	Username string `form:"username" json:"username" binding:"required"`

	jwt.StandardClaims
}
