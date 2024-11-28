package routes_client

import (
	controller_client "LearnGo/controllers/client"
	middlewares_client "LearnGo/middlewares/client"

	"github.com/gin-gonic/gin"
)

func ResultScoreRoute(r *gin.RouterGroup) {
	r.POST("/create", middlewares_client.RequireTeacher, controller_client.CreateResultScoreController)
	r.GET("/getmark", controller_client.ResultAllController)
	r.GET("/getmark/:ms", controller_client.ResultCourseController)
	r.GET("/:id", controller_client.ResultController)
	r.PATCH("/:id", middlewares_client.RequireTeacher, controller_client.ResultPatchController)
}
