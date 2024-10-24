package routes_client

import (
	controller_client "LearnGo/controllers/client"

	"github.com/gin-gonic/gin"
)

func ClassRoute(r *gin.RouterGroup) {
	r.GET("/teacher", controller_client.ClassTeacherController)
	r.GET("/student", controller_client.ClassStudentController)
	r.GET("/:id", controller_client.ClassDetailController)

}
