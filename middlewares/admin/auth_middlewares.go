package middlewares_admin

<<<<<<< HEAD
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
=======
func requireAuth() {

>>>>>>> d02509ed7fcdab80770194afdfcb89e5f7eae356
}
