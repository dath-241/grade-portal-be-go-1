package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"
	middlewares_admin "LearnGo/middlewares/admin"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup) {
	r.POST("/login", controller_admin.LoginController)
	r.POST("/create", middlewares_admin.RequireAuth, controller_admin.CreateAdminController)
}
