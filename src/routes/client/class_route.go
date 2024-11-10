package routes_client

import (
	controller_client "LearnGo/controllers/client"

	"github.com/gin-gonic/gin"
)

func ClassRoute(r *gin.RouterGroup) {
	r.GET("/account", controller_client.ClassAccountController) // lấy ra tất cả các class của account đó
	r.GET("/:id", controller_client.ClassDetailController)      // lấy ra chi tiết lớp học
	r.GET("/count/:id", controller_client.CountDocumentController)
}
