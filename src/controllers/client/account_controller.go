package controller_client

import (
	"LearnGo/helper"
	"LearnGo/models"
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func LoginController(c *gin.Context) {
	var data InterfaceAccountController
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
	collection := models.AccountModel()
	var user models.InterfaceAccount
	err = collection.FindOne(
		context.TODO(),
		bson.M{
			"email": email,
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
		"code":  "Success",
		"token": token,
		"role":  user.Role,
	})
}

func LogoutController(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{
		"code":    "Success",
		"massage": "Đăng xuất thành công",
	})
}

func AccountController(c *gin.Context) {
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

func GetInfoByIDController(c *gin.Context) {
	param := c.Param("id")
	teacher_id, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Teacher ID sai",
		})
		return
	}
	collection := models.AccountModel()
	var teacher struct {
		Name  string `bson:"name"`
		Email string `bson:"email"`
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": teacher_id, "role": "teacher"}).Decode(&teacher)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Teacher ID sai",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "success",
		"teacher": teacher,
	})
}
func CheckDuplicateOtp(ms string) bool {

	filter := bson.M{
		"ms": ms,
	}
	collection := models.OtpModel()
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false // Không tìm thấy bản ghi
	} else if err != nil {
		return false // Có lỗi khác
	}

	return true
}

func CreateOtb(c *gin.Context) {
	var otpRequest OtpRequest
	if err := c.ShouldBindJSON(&otpRequest); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	ms := otpRequest.Ms
	accCollection := models.AccountModel()
	var account models.InterfaceAccount
	err := accCollection.FindOne(context.TODO(), bson.M{"ms": ms}).Decode(&account)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Mã số không tồn tại",
		})
		return
	}
	if CheckDuplicateOtp(ms) {
		c.JSON(200, gin.H{
			"code": "error",
			"msg":  "OTP đã được gửi trước đó",
		})
		return
	}
	subject := "Xác thực mã OTP"
	otp := helper.RandomNumber(6)
	err = helper.SendMail(account.Email, subject, otp)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Gửi email thất bại",
		})
		return
	}
	otphash := helper.HashOtp(otp)
	otpCollection := models.OtpModel()
	_, err = otpCollection.InsertOne(context.TODO(), bson.M{
		"email":     account.Email,
		"ms":        ms,
		"otp":       otphash,
		"expiredAt": time.Now().Add(0), // Cập nhật đúng kiểu dữ liệu
	})
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Lưu OTP thất bại",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "success",
		"msg":  "Gửi OTP thành công",
	})
}
