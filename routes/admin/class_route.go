package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func ClassRoute(r *gin.RouterGroup) {
	r.POST("/create", controller_admin.CreateClass)
	r.PATCH("/add",controller_admin.AddStudentsToCourseHandler)
}
