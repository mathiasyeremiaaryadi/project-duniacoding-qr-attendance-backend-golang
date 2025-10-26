package routes

import (
	"qr-attendance-backend/controllers"
	"qr-attendance-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	r.GET("/api/users", middlewares.AuthMiddleware(), controllers.ListUser)
	r.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)
	r.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.CreateUser)
	r.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.CreateUser)
	r.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.CreateUser)

	return r
}
