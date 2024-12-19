package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine) {
	productGroup := router.Group("/product") // 产品模块的路由组
	{
		productGroup.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"message": "Get product info by id", "id": id})
		}) // 模拟接口
	}
}
