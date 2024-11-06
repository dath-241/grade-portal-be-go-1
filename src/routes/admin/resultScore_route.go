package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func ResultScoreRoute(r *gin.RouterGroup) {
	r.POST("/create", controller_admin.ResultScoreController)
	r.GET("/:id", controller_admin.GetResultScoreController)
}
