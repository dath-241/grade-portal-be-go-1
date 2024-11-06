package middlewares_client

import (
	"LearnGo/helper"
	"LearnGo/models"
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RequireUser(c *gin.Context) {
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
	var user models.InterfaceAccount
	collection := models.AccountModel()
	collection.FindOne(context.TODO(), bson.M{
		"_id": Claims.ID,
	}).Decode(&user)
	c.Set("user", user)
	c.Next()
}

func RequireTeacher(c *gin.Context) {
	accountGet, _ := c.Get("user")
	account := accountGet.(models.InterfaceAccount)
	if account.Role != "teacher" {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Buồn ngủ quá không code nữa đâu",
		})
		c.Abort()
		return
	}
	c.Next()
}
