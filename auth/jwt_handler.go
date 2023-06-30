package auth

import (
	"fmt"
	"gin-admin/common"
	"gin-admin/redis"
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
			if err == nil && time.Now().Unix() <= claims.ExpiresAt {
				// 如果token有效， 需要检查当前登录session
				if session, _ := redis.Get(fmt.Sprintf("Authorization:login:session:%s", claims.SessionId)); session != "" {
					// 设置uid
					c.Set("uid", claims.Id)
					c.Set("sessionId", claims.SessionId)
					// 请求继续
					c.Next()
					return
				}
			}
		}
		// 添加返回数据
		common.ResultWithError(c, common.Unauthorized)
		// 请求直接返回
		c.Abort()
	}
}
