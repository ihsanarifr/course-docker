package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	route.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"message": "hello world",
		})
	})
	route.Run(":1234")
}
