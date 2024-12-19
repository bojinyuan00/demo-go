package middleware

import (
	"demo-go/common/log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求的开始时间
		start := time.Now()

		// 获取请求信息
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		// 处理请求
		c.Next()

		// 记录请求的响应时间和状态码
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		log.AccessLogger(method, path, statusCode, latency, clientIP) // 记录访问日志
		log.SlowQueryLogger(method, path, latency)                    // 记录慢查询日志
	}
}
