package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AccountRoute(r *gin.RouterGroup) {
	r.POST("/create", controller_admin.AccountCreateController)        // thêm list các account account
	r.GET("/:id", controller_admin.AccountGetById)                     // tìm account theo id
	r.GET("/teacher", controller_admin.TeacherAccountGet)              // lấy ra tất cả account teacher
	r.GET("/student", controller_admin.StudentAccountGet)              // lấy ra tất cả account student
	r.DELETE("/delete/:id", controller_admin.DeletedAccountController) // Xóa account theo id của họ
	r.PATCH("/change/:id", controller_admin.ChangeAccountController)

}
