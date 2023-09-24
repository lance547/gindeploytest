package api

import (
	"awesomeProject/api/middleware"
	"awesomeProject/dao"
	"awesomeProject/model"
	"awesomeProject/utiles"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 注册
func Register(c *gin.Context) {
	//检验是否两个值都上传成功

	if err := c.ShouldBind(&model.User{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "verification failed",
		})
		return
	}

	username := c.PostForm("Username")
	password := c.PostForm("Password")
	//判断是否已经创建user
	flag := dao.Selectusername(username)
	if flag {
		utiles.Respfail(c, "this user already exit")
		return
	}
	//在数据库中创建数据
	dao.Adddate(username, password)
	utiles.Respsuccess(c, "user create successful")
}

// 登录
func Login(c *gin.Context) {

	//检验登录时是否上传两个值

	if err := c.ShouldBind(&model.User{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "verification failed",
		})
		return
	}
	//上传用户名和密码
	username := c.PostForm("Username")
	password := c.PostForm("Password")
	//检查用户是否存在
	flag := dao.Selectusername(username)
	if flag == false {
		utiles.Respfail(c, "this user can't found")
		return
	}
	realpassword := dao.Selectpasswordfromusername(username)
	//检查密码是否输入正确
	if realpassword != password {
		utiles.Respfail(c, "your password is wrong")
		return
	}
	//在这里生成一个pwt
	//第一步先创建一个自己的声明：
	claims := model.Myclaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 3).Unix(),
			Issuer:    "longxu",
		},
	}
	//第二步，开始生成pwt：
	//使用jwt包下的NewWithClaims函数
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//将我的秘密信息整合到tokenstring中
	tokenstring, _ := token.SignedString(middleware.Mysecret)
	//返回给前端jwt
	utiles.Respsuccess(c, tokenstring)
}
func getusernamebytoken(c *gin.Context) {
	//获取jwt中间件set的Username
	Username, _ := c.Get("Username")
	utiles.Respsuccess(c, Username.(string))

}
