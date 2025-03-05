package main

import (
	"fmt"
	"log"
	"time"

	"land_service/database"
	_ "land_service/docs"
	"land_service/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Land Service API
// @version 1.0
// @description API sử dụng Gin & Swagger
// @host localhost:9000
// @BasePath /land/api

func main() {
	database.InitDB()

	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3039"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))

	// Log tất cả request để debug
	r.Use(func(c *gin.Context) {
		log.Println("🔥 Request từ:", c.Request.Method, c.Request.RequestURI, "Origin:", c.Request.Header.Get("Origin"))
		c.Next()
	})

	// Thêm middleware CORS trước các route
	// r.Use(CORSMiddleware())

	// Ví dụ API test
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "CORS OK!"})
	})

	// Định nghĩa route
	r.POST("land/api/land", handlers.CreateLand)
	r.GET("land/api/land", handlers.GetLands)
	r.DELETE("land/api/land/:id", handlers.DeleteLand)

	// Định nghĩa route gọi API bên thứ 3
	r.GET("land/api/get-bound", handlers.GetBoundFromThirdParty)

	// Route Swagger UI
	r.GET("land/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Chạy server
	fmt.Println("🚀 Server chạy trên cổng 9000...")
	log.Fatal(r.Run(":9000"))
}
