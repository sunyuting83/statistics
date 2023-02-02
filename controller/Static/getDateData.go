package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDateData(c *gin.Context) {
	var form FormSaveData
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	d, err := MakeDateData(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "",
		"data":    d,
	})
}
