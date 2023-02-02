package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	data := gin.H{
		"status": 0,
	}
	c.JSON(http.StatusOK, data)
}
