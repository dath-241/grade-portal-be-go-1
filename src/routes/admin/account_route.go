package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AccountRoute(r *gin.RouterGroup) {
	r.POST("/create", controller_admin.AccountCreateController)
	r.GET("/:ms", controller_admin.AccountGetByMS)
	r.GET("/teacher", controller_admin.TeacherAccountGet)
	r.GET("/student", controller_admin.StudentAccountGet)
}
