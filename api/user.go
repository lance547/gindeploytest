package api

import (
	"firstproject/dao"
	"firstproject/models"
	"firstproject/utiles"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	//first to test the form whether compelete or not
	if err := c.ShouldBind(&(models.Userinformation{})); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":   500,
			"message1": "your should enter your information completely",
		})
		return
	}
	U := c.PostForm("Username") //get the username from form
	P := c.PostForm("Password") //get the password from form

	////create an  userinformation
	//myuserinformation := models.Userinformation{
	//	U,
	//	P,
	//}
	//second test for the information whether include the blank space:
	if err := dao.Testblank(U); err != nil {
		utiles.Respfail(c, "username", err)
		return
	}
	if err := dao.Testblank(P); err != nil {
		utiles.Respfail(c, "password", err)
		return
	}

	//third test for whether this user exit already before or not
	flag, err, _ := dao.Selectusername(U, dao.UUU)
	if flag {
		utiles.Respfail(c, "wrong or same name", err)
		return
	}
	//fourth to add the user information into dao
	//这里开始添加数据进入数据库
	err = dao.Addinformation(U, dao.UUU)
	if err != nil {
		utiles.Respfail(c, "add username fails", err)
		return
	}
	err = dao.Addinformation(P, dao.PPP)
	if err != nil {
		utiles.Respfail(c, "add username fails", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message4": "user create successful",
	})
}

func Login(c *gin.Context) {
	if err := c.ShouldBind(&(models.Userinformation{})); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":   500,
			"message1": "your should enter your information completely",
		})
		return
	}
	username := c.PostForm("Username")
	password := c.PostForm("Password")

	//test for this ueer whether exit aready
	flag, err, order := dao.Selectusername(username, dao.UUU)
	if !flag || err != nil || order == -1 {
		utiles.Respfail(c, "wrong or no this user", err)
		return
	}
	//test for this password whether true
	flag, err = dao.Testpassword(order, dao.PPP, password)
	if err != nil {
		utiles.Respfail(c, "wrong", err)
		return
	}
	if !flag {
		utiles.Respfail(c, "no this user", nil)
		return
	}
	//first to define our claims:
	var claims = models.Myclaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
			Issuer:    "longxu",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newtoken, err := token.SignedString(models.Mysecret)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "can't creat the token ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":         200,
		"json web token": newtoken,
	})

}

func Informationfromjwt(c *gin.Context) {
	myclaims, flag := c.Get("Username")

	if flag == false {
		c.JSON(200, gin.H{"error": "didn't get anything from token"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                 200,
		"information from token": myclaims.(string),
	})
}
