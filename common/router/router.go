package router

import (
	"demo-go/common/global"
	"demo-go/common/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() {
	fmt.Println("InitRouter == start")

	// 创建 Gin 实例
	router := gin.Default()
	// 实现跨域访问-中间件
	router.Use(middleware.Cors())
	// 加载日志中间件
	router.Use(middleware.LoggerMiddleware())

	// 按模块加载路由 start
	RegisterUserRoutes(router)    // 用户模块
	RegisterProductRoutes(router) // 产品模块
	// 按模块加载路由 end

	// 启动服务器
	port := global.Config.Server.Port
	fmt.Println("Server running on port:", port)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
