package main

import (
	"course/internal/database"
	"course/internal/exercise"
	"course/internal/middleware"
	"course/internal/user"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "root route",
		})
	})

	route.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"message": "hello world",
		})
	})

	db := database.NewDatabaseConn()
	exerciseService := exercise.NewExerciseService(db)
	userService := user.NewUserService(db)
	// exercises
	route.GET("/exercises/:id", middleware.Authentication(userService), exerciseService.GetExercise)
	route.GET("/exercises/:id/score", middleware.Authentication(userService), exerciseService.GetUserScore)

	// user
	route.POST("/register", userService.Register)
	route.POST("/login", userService.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	route.Run(":" + port)
}
