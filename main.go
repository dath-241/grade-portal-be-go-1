package main

import (
	"LearnGo/config"
	routes_admin "LearnGo/routes/admin"
	routes_client "LearnGo/routes/client"
	"fmt"
	"os"

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

	// Đăng ký các route
	routes_admin.MainRoute(app)
	routes_client.MainRoute(app)

	// Chạy server
	fmt.Println("Server đang chạy trên cổng", os.Getenv("PORT"))
	app.Run(":" + os.Getenv("PORT"))

}
