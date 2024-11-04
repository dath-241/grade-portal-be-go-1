package middlewares_admin

import (
	"LearnGo/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	// Lấy giá trị của header Authorization
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"message": "Token is required"})
		c.Abort()
		return
	}

	// Kiểm tra định dạng Bearer token
	if len(token) > 7 && strings.HasPrefix(token, "Bearer ") {
		token = token[7:] // Lấy token sau "Bearer "
	} else {
		c.JSON(401, gin.H{"message": "Invalid Authorization header"})
		c.Abort()
		return
	}
	Claims, _ := helper.ParseJWT(token)
	if Claims == nil {
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "Nguoi dung chua dang nhap",
		})
		c.Abort()
		return
	}
	c.Set("ID", Claims.ID)
	c.Next()
}
