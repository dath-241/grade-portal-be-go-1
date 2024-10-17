package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func TeacherRoute(r *gin.RouterGroup) {
	r.POST("/create", controller_admin.CreateTeacher)
}
