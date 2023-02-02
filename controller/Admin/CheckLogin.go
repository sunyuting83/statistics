package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "登陆中",
	})
}
