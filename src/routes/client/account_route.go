package routes_client

import (
	controller_client "LearnGo/controllers/client"
	middlewares_client "LearnGo/middlewares/client"

	"github.com/gin-gonic/gin"
)

func AccountRoute(r *gin.RouterGroup) {
	r.POST("/login", controller_client.LoginController)
	r.POST("/loginTele", controller_client.LoginTeleController)
	r.POST("/logout", middlewares_client.RequireUser, controller_client.LogoutController)
	r.GET("/info", middlewares_client.RequireUser, controller_client.AccountController)
	r.GET("/:id", controller_client.GetInfoByIDController)
	r.POST("/otp", controller_client.CreateOtb)
	r.POST("/resetpassword", controller_client.ResetPasswordController)
}
