package controller

import (
	"net/http"
	"statistics/database"
	"statistics/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AddAdmin(c *gin.Context) {
	var form Login
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	if len(form.UserName) < 4 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't username",
		})
		return
	}
	if len(form.Password) < 8 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't password",
		})
		return
	}

	secret_key, _ := c.Get("secret_key")
	SECRET_KEY := secret_key.(string)
	PASSWD := utils.MD5(strings.Join([]string{form.Password, SECRET_KEY}, ""))
	var admin database.Admin
	user, err := admin.CheckUserName(form.UserName)
	if err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	if len(user.Username) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "用户名已存在",
		})
		return
	}
	T := time.Now().Unix()
	admin.Username = form.UserName
	admin.Password = PASSWD
	admin.Fuck = "0"
	admin.UpdateTime = T
	admin.CreatedTime = T
	err = admin.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "添加成功",
		"data":    admin,
	})
}
