package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"
	middlewares_admin "LearnGo/middlewares/admin"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup) {
	r.POST("/login", controller_admin.LoginController)                                                                            // đăng nhập
	r.POST("/logout", middlewares_admin.RequireAuth, controller_admin.LogoutController)                                           // đăng xuất
	r.POST("/create", middlewares_admin.RequireAuth, middlewares_admin.ValidateDataAdmin, controller_admin.CreateAdminController) // thêm tài khoảng admin
	r.GET("/profile", middlewares_admin.RequireAuth, controller_admin.ProfileController)                                          // profile của bản thân
	// r.PATCH("/change")                                                                                                            // chỉnh sửa thông tin cá nhân của bản thân
}
