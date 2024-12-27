package middlewares_admin

import (
	"LearnGo/helper"
	"LearnGo/models"
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
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
	collection := models.AdminModel()
	var account models.InterfaceAdmin
	if err := collection.FindOne(context.TODO(), bson.M{
		"_id": Claims.ID,
	}).Decode(&account); err != nil {
		c.JSON(401, gin.H{
			"code": "error",
			"msg":  "Nguoi dung khong ton tai",
		})
		c.Abort()
		return
	}
	c.Set("ID", Claims.ID)
	c.Next()
}
