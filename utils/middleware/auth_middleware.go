package middleware

import (
	"net/http"
	jwtHelper "pro05shopping/utils/jwt"

	"github.com/gin-gonic/gin"
)

// 管理员授权
func AuthAdminMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil && decodedClaims.IsAdmin {
				c.Next()
				c.Abort()
				return
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限访问!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "你没有授权!"})
		}
		c.Abort()
		return
	}
}

// 用户授权
func AuthUserMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				c.Set("userId", decodedClaims.UserId)
				c.Next()
				c.Abort()
				return
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限访问!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "你没授权!"})
		}
		c.Abort()
		return
	}
}
