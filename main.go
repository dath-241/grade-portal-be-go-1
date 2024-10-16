package main

import (
	"LearnGo/config"
	routes_admin "LearnGo/routes/admin"
	routes_client "LearnGo/routes/client"
	"fmt"
	"net/http"
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
	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Nếu là yêu cầu OPTIONS, trả về 200 OK và dừng lại
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Tiếp tục với request tiếp theo
		c.Next()
	})

	// Đăng ký các route
	routes_admin.MainRoute(app)
	routes_client.MainRoute(app)

	// Chạy server
	fmt.Println("Server đang chạy trên cổng", os.Getenv("PORT"))
	app.Run(":" + os.Getenv("PORT"))
}
