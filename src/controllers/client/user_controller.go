package controller_client

import (
	"LearnGo/helper"
	"LearnGo/models"
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func LoginController(c *gin.Context) {
	var data InterfaceUserController
	// lấy dữ liệu từ front end
	c.BindJSON(&data)
	payload, err := idtoken.Validate(context.Background(), data.IDToken, os.Getenv("YOUR_CLIENT_ID"))
	if err != nil {
		fmt.Println("Khong co token:", err)
		c.JSON(401, gin.H{"error": "Token khong hop le"})
		return
	}
	// lay ra email
	email, emailOk := payload.Claims["email"].(string)
	if !emailOk {
		c.JSON(400, gin.H{"error": "khong lay duoc thong tin nguoi dung"})
		return
	}
	fmt.Println(email, data.Role)
	// tim kiem nguoi dung da co trong db khong
	collection := models.UserModel()
	var user models.InterfaceUser
	err = collection.FindOne(
		context.TODO(),
		bson.M{
			"email": email,
			"role":  data.Role,
		},
	).Decode(&user)
	fmt.Println(user)
	if err != nil {
		c.JSON(400, gin.H{"error": "khong lay duoc thong tin nguoi dung trong dữ liệu vui lòng liên hệ admin để thêm bạn vào"})
		return
	}
	token := helper.CreateJWT(user.ID)
	c.SetCookie("token", token, 3600*24, "/", "", false, true)
	c.JSON(200, gin.H{
		"code": "Success",
	})
}

func LogoutController(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{
		"code":    "Success",
		"massage": "Đăng xuất thành công",
	})
}

func UserController(c *gin.Context) {
	user, _ := c.Get("user")
	if user == "" {
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "Không có người dùng",
		})
	}
	c.JSON(200, gin.H{
		"code": "success",
		"user": user,
	})
}
