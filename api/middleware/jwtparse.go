package middleware

import (
	"errors"
	"firstproject/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// use jwt middleware to parse the token
func JWTparse() func(*gin.Context) {
	return func(c *gin.Context) {
		//through the request to get the header'key :Artherization.
		arth := c.Request.Header.Get("Artherization")
		if arth == "" {
			c.JSON(http.StatusOK, gin.H{
				"status":  500,
				"message": "Artherization is nil",
			})
			c.Abort()
			return
		}
		arthsplit := strings.SplitN(arth, " ", 2)
		if arthsplit[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"Status":  500,
				"message": "the format of Artherization is wrong",
			})
			c.Abort()
			return
		}
		token := arthsplit[1]
		//import the Parsejwt function to parse the token:
		message, err := Parsejwt(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  500,
				"message": err,
			})
			c.Abort()
			return

		}
		//attention must through the struct to get the username information ,if you use the interface that you
		//could't get the information you want !!!
		c.Set("Username", message.Username) //set this claim that I can get this claim later.
		c.Next()

	}
}
func Parsejwt(token string) (*models.Myclaims, error) {
	t, e := jwt.ParseWithClaims(token, &models.Myclaims{}, func(token *jwt.Token) (interface{}, error) {
		return models.Mysecret, nil
	})
	if e != nil {
		return nil, e
	}
	if token, ok := t.Claims.(*models.Myclaims); ok && t.Valid {
		return token, nil
	}
	return nil, errors.New("invaild token")

}
