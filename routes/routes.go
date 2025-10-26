package routes

import (
	"qr-attendance-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/register", controllers.Register)

	return r
}
