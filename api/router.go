package api

import (
	"firstproject/api/middleware"
	"github.com/gin-gonic/gin"
)

//router包用于存放使用的路由

func Apirouter() {
	//set an engine:
	r := gin.Default()
	//use middleware here:
	r.Use(middleware.CORS()) //use the CORS middleware

	//http response here:
	r.POST("/register", Register) //sent a register request
	r.POST("/login", Login)       //sent a login request
	jwtrouter := r.Group("/user")
	{
		jwtrouter.Use(middleware.JWTparse())      //use the JWTparse middleware
		jwtrouter.GET("/get", Informationfromjwt) //set a get information from jwt request
	}
	//let this engine runing
	r.Run()
}
