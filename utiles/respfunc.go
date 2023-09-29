package utiles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// response successful
func Respsuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
	})
}

// response failed
func Respfail(c *gin.Context, message string, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status":  500,
		"message": message,
		"error":   err,
	})
}
