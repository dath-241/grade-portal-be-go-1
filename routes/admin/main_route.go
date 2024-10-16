package routes_admin

import "github.com/gin-gonic/gin"

func MainRoute(r *gin.Engine) {
	AuthRoute(r.Group("/admin/api"))
}
