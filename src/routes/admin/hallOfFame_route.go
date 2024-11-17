package routes_admin

import (
	controller_admin "LearnGo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func HallOfFameRoute(r *gin.RouterGroup) {
	r.POST("/update", controller_admin.CreateHallOfFame)
}
