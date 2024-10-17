package controller_admin

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
	var data AuthController
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
<<<<<<< HEAD
=======
	fmt.Println(email)
>>>>>>> d02509ed7fcdab80770194afdfcb89e5f7eae356
	// tim kiem nguoi dung da co trong db khong
	collection := models.AdminModel()
	var user models.InterfaceAdmin
	collection.FindOne(
		context.TODO(),
		bson.M{
			"email": email,
		},
	).Decode(&user)
<<<<<<< HEAD
=======
	fmt.Println(user)
>>>>>>> d02509ed7fcdab80770194afdfcb89e5f7eae356
	token := helper.CreateJWT(email)
	c.SetCookie("token", token, 3600*24, "/", "", false, true)
	c.JSON(200, gin.H{
		"code": "Success",
	})
}
