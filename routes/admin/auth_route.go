package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup) {
	r.POST("/login", controller_admin.LoginController)
}
