package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func CourseRoute(r *gin.RouterGroup) {
	r.POST("create", controller_admin.CreateCourse)
	// r.GET("create", controller_admin.Course)
}
