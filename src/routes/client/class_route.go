package routes_client

import (
	controller_client "LearnGo/controllers/client"

	"github.com/gin-gonic/gin"
)

func ClassRoute(r *gin.RouterGroup) {
	r.GET("/account", controller_client.ClassAccountController)
	r.GET("/:id", controller_client.ClassDetailController)
}
