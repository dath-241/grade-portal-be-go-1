package middlewares_client

import (
	"LearnGo/helper"
	"LearnGo/models"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RequireUser(c *gin.Context) {
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
	var user models.InterfaceUser
	collection := models.UserModel()
	collection.FindOne(context.TODO(), bson.M{
		"_id": Claims.ID,
	}).Decode(&user)
	c.Set("user", user)
	c.Next()
}
