package router

import (
	"demo-go/app/provider"
	"demo-go/common/global"
	"demo-go/common/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func InitRouter() {
	//add router initialization code here
	fmt.Println("Router initialized")

	// 创建 Gin 实例
	router := gin.Default()

	// 实现跨域访问
	router.Use(middleware.Cors())

	// 加载日志中间件
	router.Use(middleware.LoggerMiddleware())

	userController, _ := provider.InitializeUserController()
	router.GET("/user/test/:id", userController.GetUserInfo) // 定义用户) // 获取用户信息

	// 路由映射
	router.GET("/user/:user_id", func(c *gin.Context) {
		user_id := c.Param("user_id")
		time.Sleep(600 * time.Millisecond) // 模拟慢请求
		fmt.Println("Get user info by user_id:", user_id)
	}) // 定义用户) // 获取用户信息

	// 启动服务器
	port := global.Config.Server.Port
	fmt.Println("Server running on port:", port)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
