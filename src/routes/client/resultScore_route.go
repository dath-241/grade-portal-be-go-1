package routes_client

import (
	controller_client "LearnGo/controllers/client"

	"github.com/gin-gonic/gin"
)

func ResultScoreRoute(r *gin.RouterGroup) {
	r.POST("/create", controller_client.CreateResultScoreController)
	r.GET("/:id", controller_client.ResultController)
	r.PATCH("/:id", controller_client.ResultPatchController)
}
