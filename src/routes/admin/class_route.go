package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func ClassRoute(r *gin.RouterGroup) {
	r.POST("/create", controller_admin.CreateClass)
	r.GET("/:id", controller_admin.GetClassByClassID)
	r.GET("/account/:id", controller_admin.GetAllClassesByAccountID)

}
