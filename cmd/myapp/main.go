package main

import (
	"goject/config"
	"goject/routes"

	"github.com/gin-gonic/gin"

	_ "goject/docs" // import local generated docs package

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GoJect API
// @version 1.0

// @host localhost:8080
// @BasePath /api/v1
func main() {
	// อ่าน config
	config.LoadConfig()

	// สร้าง Gin engine
	router := gin.Default()

	// ตั้งค่า middleware
	// router.Use(myappMiddleware)

	// ตั้งค่าเส้นทาง
	routes.SetupRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// เริ่มเซิร์ฟเวอร์
	router.Run(":8080")
}
