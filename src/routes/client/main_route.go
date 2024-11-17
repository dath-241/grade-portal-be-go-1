package routes_client

import (
	middlewares_client "LearnGo/middlewares/client"

	"github.com/gin-gonic/gin"
)

func MainRoute(r *gin.Engine) {
	HomeRouter(r.Group("/"))
	AccountRoute(r.Group("/api"))
	protectedGroup := r.Group("/api")
	protectedGroup.Use(middlewares_client.RequireUser)
	HallOfFameRoute(protectedGroup.Group("/HOF"))
	ClassRoute(protectedGroup.Group("/class"))
	CourseRoute(protectedGroup.Group("/course"))
	ResultScoreRoute(protectedGroup.Group("/resultScore"))
}
