package main

import (
	"course/internal/database"
	"course/internal/exercise"
	"course/internal/middleware"
	"course/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	db := database.NewConnDatabase()
	exerciseService := exercise.NewExerciseService(db)
	userService := user.NewUserService(db)
	r.GET("/exercises/:id", middleware.JWTAuth(userService), exerciseService.GetExerciseByID)
	r.GET("/exercises/:id/score", middleware.JWTAuth(userService), exerciseService.GetUserScore)

	r.POST("/register", userService.Register)
	r.POST("/login", userService.Login)
	r.Run(":1234")
}
