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
	// Để query tìm teacher/student theo ms
	// API request theo định dạng
	// /admin/api/account/student?ms=ID
	// Thay ID bằng ms để tra cứu
}
