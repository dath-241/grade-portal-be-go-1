package routes_client

import (
	controller_client "LearnGo/controllers/client"

	"github.com/gin-gonic/gin"
)

func HomeRouter(r *gin.RouterGroup) {
	r.GET("/", controller_client.HomeController)
}
