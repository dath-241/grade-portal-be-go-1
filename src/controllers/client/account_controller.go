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
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Data không nhận được",
		})
		return
	}
	payload, err := idtoken.Validate(context.Background(), data.IDToken, os.Getenv("YOUR_CLIENT_ID"))
	if err != nil {
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
		"code": "Success",
		"msg":  "Đăng xuất thành công",
	})
}

func AccountController(c *gin.Context) {
	user, _ := c.Get("user")
	if user == "" {
		c.JSON(401, gin.H{
			"code": "error",
			"msg":  "Không có người dùng",
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
	fmt.Println(err)
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
		"expiredAt": time.Now().Add(5 * 60 * 1000),
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

func ResetPasswordController(c *gin.Context) {
	var resgister RegisterInterface
	if err := c.ShouldBindJSON(&resgister); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	var account models.InterfaceAccount
	collection_account := models.AccountModel()
	err := collection_account.FindOne(context.TODO(), bson.M{"ms": resgister.Ms}).Decode(&account)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	var otp models.InterfaceOtp
	collection_otp := models.OtpModel()
	err = collection_otp.FindOne(context.TODO(), bson.M{
		"email": account.Email,
		"ms":    resgister.Ms,
		"otp":   helper.HashOtp(resgister.Otp),
	}).Decode(&otp)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	_, err = collection_otp.DeleteOne(context.TODO(), bson.M{
		"email": account.Email,
		"ms":    resgister.Ms,
		"otp":   helper.HashOtp(resgister.Otp),
	})
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	_, err = collection_account.UpdateOne(context.TODO(), bson.M{
		"_id": account.ID,
	}, bson.M{
		"$set": bson.M{
			"password": helper.HashOtp(resgister.Password),
		},
	})
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "success",
		"msg":  "Thay đổi password thành công",
	})
}

func LoginTeleController(c *gin.Context) {
	var login LoginInterface
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	collection := models.AccountModel()
	var account models.InterfaceAccountTelegram
	err := collection.FindOne(context.TODO(), bson.M{
		"ms": login.Ms,
	}).Decode(&account)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	if helper.HashOtp(login.Password) != account.Password {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	// semenster := helper.Set_semester()
	var classAccount []models.InterfaceClass
	collection_class := models.ClassModel()
	cursor_class, err := collection_class.Find(context.TODO(), bson.M{
		"listStudent_ms": account.Ms,
		// "semester": bson.M{
		// 	"$in": [2]string{semenster.PREV, semenster.CUREENT},
		// },
	})
	if err != nil {
		c.JSON(401, gin.H{
			"code": "error",
			"msg":  "1",
		})
		return
	}
	defer cursor_class.Close(context.TODO())
	if err := cursor_class.All(context.TODO(), &classAccount); err != nil {
		c.JSON(401, gin.H{
			"code": "error",
			"msg":  "2",
		})
		return
	}
	var IDs []bson.ObjectID
	checkMap := make(map[bson.ObjectID]string)
	for _, item := range classAccount {
		IDs = append(IDs, item.CourseId)
		checkMap[item.CourseId] = item.Semester
	}
	var listCourse []models.InterfaceCourse
	collection_course := models.CourseModel()
	cursor_course, err := collection_course.Find(context.TODO(), bson.M{
		"_id": bson.M{
			"$in": IDs,
		},
	})
	if err != nil {
		c.JSON(401, gin.H{
			"code": "error",
			"msg":  "3",
		})
		return
	}
	defer cursor_course.Close(context.TODO())
	if err := cursor_course.All(context.TODO(), &listCourse); err != nil {
		c.JSON(401, gin.H{
			"code": "error",
			"msg":  "4",
		})
		return
	}
	var msList []string
	for _, item := range listCourse {
		msList = append(msList, item.MS+"-"+checkMap[item.ID])
	}
	token := helper.CreateJWT(account.ID)
	c.JSON(200, gin.H{
		"code":       "success",
		"token":      token,
		"listCourse": msList,
		"role":       account.Role,
	})
}
