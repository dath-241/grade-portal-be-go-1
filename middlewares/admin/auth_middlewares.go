package middlewares_admin

import (
	"LearnGo/helper"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	token, _ := c.Cookie("token")
	Claims, _ := helper.ParseJWT(token)
	if Claims == nil {
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "Nguoi dung chua dang nhap",
		})
		c.Abort()
		return
	}
	c.Next()
}
