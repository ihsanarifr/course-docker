package route

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine) {
	r.GET("/abc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "abc",
		})
	})
}
