package models

import "github.com/dgrijalva/jwt-go"

// jwt's claims
type Myclaims struct {
	Username string `form:"Username" json:"Username" binding:"required"`
	jwt.StandardClaims
}

var Mysecret = []byte("secret")

// user's information
type Userinformation struct {
	Username string `form:"Username" json:"Username"binding:"required"`
	Password string `form:"Password"json:"Password"binding:"required"`
}
