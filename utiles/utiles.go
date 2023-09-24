package utiles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//封装工具函数，减少代码重复利用

func Respsuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
	})
}
func Respfail(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  500,
		"message": message,
	})
}
