package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func ClassRoute(r *gin.RouterGroup) {
	r.POST("/create", controller_admin.CreateClass)                  // tạo 1 lớp học mới
	r.GET("/:id", controller_admin.GetClassByClassID)                // lấy ra chi tiết lớp học
	r.GET("/account/:id", controller_admin.GetAllClassesByAccountID) // lấy ra tất cả lớp học của id account đó
	r.GET("/course/:id", controller_admin.GetClassByCourseID)        // lấy ra tất cả lớp học của id course đó
	r.PATCH("/add", controller_admin.AddStudentsToCourseHandler)     //thêm học sinh vào lớp học đó
	r.DELETE("/delete/:id", controller_admin.DeleteClassController)  // xóa lớp học theo id lớp học
	r.PATCH("/change/:id", controller_admin.ChangeClassController)   // chinhr sửa thông tin lớp học

}
