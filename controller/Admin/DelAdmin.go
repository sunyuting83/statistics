package controller

import (
	"net/http"
	"statistics/database"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID int64 `form:"id" json:"id" xml:"id"  binding:"required"`
}

func DeleteAdmin(c *gin.Context) {
	var form User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	var admin database.Admin
	user, err := admin.CheckID(form.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	admin.DeleteOne(form.ID)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "成功删除管理员",
		"id":      user.ID,
	})
}
