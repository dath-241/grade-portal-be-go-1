package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func CourseRoute(r *gin.RouterGroup) {
	r.POST("/create", controller_admin.CreateCourse)       // tạo khóa học
	r.GET("/:id", controller_admin.GetCourseByCourseID)    // lấy chi tiết khóa học
	r.GET("/all", controller_admin.GetAllCourseController) // lấy ra tất cả khóa học
	r.DELETE("/delete/:id", controller_admin.DeleteCourseController)
	r.PATCH("/change/:id", controller_admin.ChangeCourseController)
	// làm query cho tất cả khóa học này là tìm theo ms cho get all trên
}
