package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/transport/http/response"
	"github.com/zhang/bms/internal/utils"
)

func IgnorePath(path string, ignorePaths []string) bool {
	for _, ignorePath := range ignorePaths {
		if path == ignorePath {
			return true
		}
	}
	return false
}

func JwtAuthMiddleware(ignorePaths *[]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断是否需要进行登入
		path := c.Request.URL.Path
		if IgnorePath(path, *ignorePaths) {
			c.Next()
			return
		}

		// 1.从 gin.Context 中获取 token, 在这里通header中的 Authorization 中, 并使用Bearer开头
		token := c.Request.Header.Get("Authorization")
		// 2.判断请求头中是否携带 Authorization
		if token == "" {
			fmt.Println("未携带请求头")
			response.FailWithMsg(
				c, http.StatusNonAuthoritativeInfo,
				"请求头中的authorization为空",
			)
			c.Abort()
			return
		}

		// 3.请求头中携带了 authorization, 判断格式是否正确
		parts := strings.Split(token, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			fmt.Println("请求头中的auth格式有误")
			response.FailWithMsg(
				c, http.StatusBadRequest, "请求头中的auth格式有误",
			)
			c.Abort()
			return
		}

		// 4.解析token
		mc, err := utils.ParseJwt(parts[1])
		if err != nil {
			fmt.Println("无效的token")
			response.FailWithMsg(c, http.StatusBadRequest, "无效的Token")
			c.Abort()
			return
		}

		c.Set("username", mc.Username)
		c.Next()

	}
}
