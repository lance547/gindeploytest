package middleware

import (
	"awesomeProject/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var Mysecret = []byte("this is a secret!!!!")

// 基于jwt的认证中间件
// 使用jwtmiddleware函数来作为中间件
func JWTmiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//客户端携带token有三种实现方式，请求头，请求体，请求尾，
		//这里我们假设token放在header的Arthorization中，并且使用Bearer开头
		//具体的实现方式还是由具体业务决定
		arthheader := c.Request.Header.Get("Artherization")
		//判断获取的请求头是否为空
		if arthheader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    "2003",
				"message": "请求头中的arth为空",
			})
			c.Abort()
			return
		}
		//进行按空格分割
		//请求头中由bearer和token组成，他们由空格分开，所以直接从空格切断，，，以下为splitN方法注释:
		// SplitAfterN slices s into substrings after each instance of sep and
		// returns a slice of those substrings.
		//
		// The count determines the number of substrings to return:
		//
		//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
		//	n == 0: the result is nil (zero substrings)
		//	n < 0: all substrings
		//
		parts := strings.SplitN(arthheader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code":    2004,
				"message": "arth的格式有误",
			})
		}
		//开始提取请求头中的token信息
		//使用解析函数来解析jwt
		MYCLAIMS, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    2005,
				"message": err,
			})
			c.Abort()
			return
		}
		//将解析的要求保存
		c.Set("Username", MYCLAIMS.Username)
		// Next should be used only inside middleware.
		// It executes the pending handlers in the chain inside the calling handler.
		// See example in GitHub.
		c.Next()
	}
}

// 解析jwt前获得的是string类型的字符串，
func ParseToken(tokenstring string) (*model.Myclaims, error) {
	//根据要求解析token字符串：
	token, err := jwt.ParseWithClaims(tokenstring, &model.Myclaims{}, func(token *jwt.Token) (any, error) {
		return Mysecret, nil
	})

	if err != nil {
		return nil, err
	}
	//这里有点不太懂，明天再重点看看：
	claims, ok := token.Claims.(*model.Myclaims)
	//检验token是否可用，可用返回，不可用退出
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invaild token")
}
