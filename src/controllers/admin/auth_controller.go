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
	// tim kiem nguoi dung da co trong db khong
	collection := models.AdminModel()
	var user models.InterfaceAdmin
	err = collection.FindOne(
		context.TODO(),
		bson.M{
			"email": email,
		},
	).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "khong lay duoc thong tin nguoi dung trong dữ liệu vui lòng liên hệ admin để thêm bạn vào"})
		return
	}
	token := helper.CreateJWT(user.ID)
	c.SetCookie("token", token, 3600*24, "http://127.0.0.1:5500", "http://127.0.0.1:5500/src/index.html", false, true)
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

func CreateAdminController(c *gin.Context) {
	adminData, _ := c.Get("adminData")
	data := adminData.(InterfaceAdminController)
	collection := models.AdminModel()
	createBy, _ := c.Get("ID")
	var Admin models.InterfaceAdmin
	err := collection.FindOne(
		context.TODO(),
		bson.M{
			"$or": bson.A{
				bson.M{"email": data.Email},
				bson.M{"ms": data.Ms},
			},
		},
	).Decode(&Admin)
	if err == nil {
		c.JSON(400, gin.H{
			"code":    "error",
			"massage": "Bảng ghi của admin này đã được lưu trong database trước đó",
		})
		return
	}
	collection.InsertOne(context.TODO(), bson.M{
		"email":     data.Email,
		"name":      data.Name,
		"faculty":   data.Faculty,
		"ms":        data.Ms,
		"createdBy": createBy,
	})
	c.JSON(201, gin.H{
		"code": "vao duoc trang createAdmin",
	})
}
