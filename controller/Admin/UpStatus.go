package controller

import (
	"fmt"
	"net/http"
	"statistics/database"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpStatusAdmin(c *gin.Context) {
	var form User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	var (
		admin   database.Admin
		Fuck    string = "1"
		FuckStr string = "锁定"
	)
	user, err := admin.CheckID(form.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	if user.Fuck == "1" {
		fmt.Println(user.Fuck)
		Fuck = "0"
		FuckStr = "解锁"
	}
	admin.Fuck = Fuck
	a, err := admin.ResetPassword(user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
			"user":    a.Fuck,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": strings.Join([]string{"成功", FuckStr, "管理员"}, ""),
		"user":    a.Fuck,
	})
}
