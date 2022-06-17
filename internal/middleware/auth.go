package middleware

import (
	"context"
	"course/internal/user"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	bearer = "Bearer "
)

func JWTAuth(us *user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"message": "anda tidak diizinkan mengakses ini",
			})
			c.Abort()
		}
		if !strings.HasPrefix(authHeader, bearer) {
			c.JSON(401, gin.H{
				"message": "anda tidak diizinkan mengakses ini",
			})
			c.Abort()
		}
		auths := strings.Split(authHeader, " ")
		data, err := us.DecriptJWT(auths[1])
		if err != nil {
			c.JSON(401, gin.H{
				"message": "anda tidak diizinkan mengakses ini",
			})
			c.Abort()
		}
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", data["user_id"])
		c.Request = c.Request.WithContext(ctxUserID)
		c.Next()
	}
}
