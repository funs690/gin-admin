package auth

import (
	"gin-admin/common"
	"github.com/gin-gonic/gin"
	"time"
)

// 设置权限验证
func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.GetHeader("token")
		// token 验证
		if token != "" {
			// 解析token
			claims, err := ParseToken(token)
			if err == nil || time.Now().Unix() > claims.ExpiresAt {
				// 设置uid
				c.Set("uid", claims.Id)
				// 请求继续
				c.Next()
				return
			}
		}
		// 添加返回数据
		common.ResultWithError(c, common.Unauthorized)
		// 请求直接返回
		c.Abort()
	}
}
