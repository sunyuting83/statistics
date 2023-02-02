package router

import (
	controller "statistics/controller"
	Admin "statistics/controller/Admin"
	Static "statistics/controller/Static"
	utils "statistics/utils"

	"github.com/gin-gonic/gin"
)

// SetConfigMiddleWare set config
func SetConfigMiddleWare(SECRET_KEY, CODE string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("secret_key", SECRET_KEY)
		c.Set("code", CODE)
		c.Writer.Status()
	}
}

// InitRouter make router
func InitRouter(SECRET_KEY, CODE string) *gin.Engine {
	router := gin.Default()
	router.Use(utils.CORSMiddleware())
	api := router.Group("/api")
	api.Use(SetConfigMiddleWare(SECRET_KEY, CODE))
	{
		router.GET("/", controller.Index)
		api.POST("/addadmin", utils.VerifyMiddleware(), Admin.AddAdmin)
		api.PUT("/repassword", utils.VerifyMiddleware(), Admin.ResetPassword)
		api.DELETE("/deladmin", utils.VerifyMiddleware(), Admin.DeleteAdmin)
		api.GET("/checklogin", utils.VerifyMiddleware(), Admin.CheckLogin)
		api.GET("/adminlist", utils.VerifyMiddleware(), Admin.AdminList)
		api.PUT("/upstatus", utils.VerifyMiddleware(), Admin.UpStatusAdmin)
		api.POST("/loginadmin", Admin.Sgin)
		api.GET("/putit", Static.Putit)
	}

	return router
}
