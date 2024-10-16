package routes_client

import "github.com/gin-gonic/gin"

func MainRoute(r *gin.Engine) {
	HomeRouter(r.Group("/api"))
}
