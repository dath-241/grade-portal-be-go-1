package routes_admin

import (
	"github.com/gin-gonic/gin"

	controller_admin "LearnGo/controllers/admin"
)

func HallOfFameRoute(r *gin.RouterGroup) {
	r.POST("/update", controller_admin.CreateHallOfFame)
	r.GET("/all", controller_admin.GetPrevSemesterAllHallOfFame)
}
