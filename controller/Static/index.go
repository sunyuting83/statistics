package controller

import (
	"net/http"
	"statistics/database"
	"statistics/utils"

	"github.com/gin-gonic/gin"
)

func GetIndexData(c *gin.Context) {
	var hostList *database.Host
	List, err := hostList.GetAllHost()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	start, end := utils.GetDateTime()

	var today *database.Static
	Today, err := today.GetTodayStatic(start, end)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	Count := int64(0)
	ComputNumber := len(Today)

	for _, item := range Today {
		Count += item.Count
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   0,
		"message":  "",
		"hostlist": List,
		"count":    Count,
		"comput":   ComputNumber,
	})
}
