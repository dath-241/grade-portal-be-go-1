package routes_client

import (
	controller_client "LearnGo/controllers/client"

	"github.com/gin-gonic/gin"
)

func HallOfFameRoute(r *gin.RouterGroup) {
	r.GET("/all", controller_client.GetPrevSemesterAllHallOfFame)
}
