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
	ClassRoute(protectedGroup.Group("/class"))
	ResultScoreRoute(protectedGroup.Group("/resultScore"))
}
