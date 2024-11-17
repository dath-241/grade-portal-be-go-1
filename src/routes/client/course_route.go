package routes_client

import (
	controller_client "LearnGo/controllers/client"

	"github.com/gin-gonic/gin"
)

func CourseRoute(r *gin.RouterGroup) {
	r.GET("/:id", controller_client.GetCourseByCourseID) // lấy ra chi tiết khoa học
}
