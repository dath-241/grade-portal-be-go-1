package main

import (
	"LearnGo/config"
	routes_admin "LearnGo/routes/admin"
	routes_client "LearnGo/routes/client"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Tải biến môi trường
	godotenv.Load()
	config.ConnectMongoDB(os.Getenv("MONGO_URL"))

	// Tạo một instance của Gin
	app := gin.Default()

	// Cấu hình CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://test-jcz3.vercel.app", "http://localhost:3000"}, // Chỉ cho phép origin cụ thể
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Đăng ký các route
	routes_admin.MainRoute(app)
	routes_client.MainRoute(app)

	// Chạy server
	fmt.Println("Server đang chạy trên cổng", os.Getenv("PORT"))
	app.Run(":" + os.Getenv("PORT"))

}
