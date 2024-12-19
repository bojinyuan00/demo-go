package router

import (
	"demo-go/app/provider"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user") // 用户模块的路由组
	{
		userController, _ := provider.InitializeUserController() //实例化userController控制器（容器依赖注入）
		userGroup.GET("/test/:id", userController.GetUserInfo)   // 获取用户信息
		userGroup.GET("/:user_id", func(c *gin.Context) {
			userID := c.Param("user_id")
			c.JSON(200, gin.H{"message": "Get user info by user_id", "user_id": userID})
		}) // 模拟接口
	}
}
