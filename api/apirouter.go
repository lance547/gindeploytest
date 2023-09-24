package api

import (
	"awesomeProject/api/middleware"
	"github.com/gin-gonic/gin"
)

// 接口层，封装详细的逻辑实现及路由
func Router() {
	r := gin.Default()
	//使用cors中间件
	r.Use(middleware.CORS())
	//分组路由
	Userrouter := r.Group("/user")
	{
		//使用解析jwt的中间件
		Userrouter.Use(middleware.JWTmiddleware())
		//获取解析出来的jwt中的claims
		Userrouter.GET("/Get", getusernamebytoken)
	}

	//http请求
	r.POST("/register", Register)
	r.POST("/login", Login)
	//运行
	r.Run()
}
