package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/pkg/logger"
	"go.uber.org/zap"
)

// GinLogger 接收 Gin 框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		cost := time.Since(start)

		var msg string
		if v, exists := c.Get("msg"); exists {
			msg, _ = v.(string) // 强制转为 string（如果不是 string，会是空字符串）
		}

		// 记录请求日志
		logger.Log.Info("request",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("error", msg),
			zap.Duration("cost", cost),
		)
	}
}
