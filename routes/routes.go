package routes

import (
	"goject/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// กำหนดเส้นทาง
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", controllers.GetUsers)
		v1.POST("/users", controllers.CreateUser)
	}
}
