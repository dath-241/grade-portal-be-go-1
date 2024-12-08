package controller_admin

import (
	"LearnGo/models"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func AccountCreateController(c *gin.Context) {
	var newUsers []InterfaceAccountController
	// Bind JSON từ body của request vào struct
	if err := c.ShouldBindJSON(&newUsers); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Data không nhận được",
		})
		return
	}
	userCollection := models.AccountModel()

	// Lấy tất cả tài khoản từ cơ sở dữ liệu
	var existingUsers []models.InterfaceAccount
	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users from database"})
		return
	}
	if err := cursor.All(context.TODO(), &existingUsers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding users"})
		return
	}
	CreatedBy, _ := c.Get("ID")

	m := make(map[string]bool)
	var filterAccount []InterfaceAccountController
	var errorAccount []InterfaceAccountController
	for _, account := range existingUsers {
		m[account.Email] = true
		m[account.Ms] = true
	}
	for _, newAccount := range newUsers {
		if !m[newAccount.Email] && !m[newAccount.Ms] && strings.HasSuffix(newAccount.Email, "@hcmut.edu.vn") && (newAccount.Role == "student" || newAccount.Role == "teacher") {
			newAccount.CreatedBy = CreatedBy
			newAccount.ExpiredAt = time.Now().AddDate(5, 0, 0)
			filterAccount = append(filterAccount, newAccount)
		} else {
			errorAccount = append(errorAccount, newAccount)
		}
	}
	// for _, newUser := range newUsers {

	// 	if !strings.HasSuffix(newUser.Email, "@hcmut.edu.vn") {
	// 		invalidEmails = append(invalidEmails, newUser)
	// 		continue
	// 	}
	// Chèn các tài khoản hợp lệ vào cơ sở dữ liệu
	if len(filterAccount) > 0 {
		_, err := userCollection.InsertMany(context.TODO(), filterAccount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating valid accounts"})
			return
		}
	}
	// Trả về phản hồi, thông báo người dùng nào đã được thêm và ai bị trùng lặp
	c.JSON(200, gin.H{
		"code":          "success",
		"errorAccount":  errorAccount,
		"accessAccount": filterAccount,
	})
}

func AccountGetById(c *gin.Context) {
	param := c.Param("id") // Lấy giá trị "" từ URL

	accountId, err := bson.ObjectIDFromHex(param)

	if err != nil {
		c.JSON(400, gin.H{
			"code": err,
			"msg":  "Lỗi id gửi lên",
		})
		return
	}

	userCollection := models.AccountModel()

	// Tạo biến để lưu kết quả
	var user models.InterfaceAccount

	// Tìm trong MongoDB theo trường MS
	err = userCollection.FindOne(context.TODO(), bson.M{"_id": accountId}).Decode(&user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			// Nếu không tìm thấy user
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// Xử lý lỗi khác
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
		return
	}

	// Trả về thông tin user
	c.JSON(http.StatusOK, gin.H{
		"status":  "User found successfully",
		"account": user,
	})
}

func TeacherAccountGet(c *gin.Context) {
	userCollection := models.AccountModel()
	query := c.Query("ms")

	if query == "" {
		// Tạo biến lưu kết quả
		var users []models.InterfaceAccount

		cursor, err := userCollection.Find(context.TODO(), bson.M{"role": "teacher"})
		// Xử lý lỗi
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				// Nếu không tìm thấy user
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			// Xử lý lỗi khác
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
			return
		}

		if err := cursor.All(context.TODO(), &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
			return
		}

		// Trả về thông tin users
		c.JSON(http.StatusOK, gin.H{
			"status":      "Users found successfully",
			"foundedUser": users,
		})
	} else {
		// Tạo biến để lưu kết quả
		var user models.InterfaceAccount

		// Tìm trong MongoDB theo trường MS
		err := userCollection.FindOne(context.TODO(), bson.M{"ms": query}).Decode(&user)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				// Nếu không tìm thấy user
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			// Xử lý lỗi khác
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
			return
		}

		// Trả về thông tin user
		c.JSON(http.StatusOK, gin.H{
			"status":      "User found successfully",
			"foundedUser": user,
		})
	}
}

func StudentAccountGet(c *gin.Context) {
	userCollection := models.AccountModel()
	query := c.Query("ms")

	if query == "" {
		// Tạo biến lưu kết quả
		var users []models.InterfaceAccount

		cursor, err := userCollection.Find(context.TODO(), bson.M{"role": "student"})
		// Xử lý lỗi
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				// Nếu không tìm thấy user
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			// Xử lý lỗi khác
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
			return
		}

		if err := cursor.All(context.TODO(), &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
			return
		}

		// Trả về thông tin users
		c.JSON(http.StatusOK, gin.H{
			"status":      "Users found successfully",
			"foundedUser": users,
		})
	} else {
		// Tạo biến để lưu kết quả
		var user models.InterfaceAccount

		// Tìm trong MongoDB theo trường MS
		err := userCollection.FindOne(context.TODO(), bson.M{"ms": query}).Decode(&user)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				// Nếu không tìm thấy user
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			// Xử lý lỗi khác
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
			return
		}

		// Trả về thông tin user
		c.JSON(http.StatusOK, gin.H{
			"status":      "User found successfully",
			"foundedUser": user,
		})
	}
}

func DeletedAccountController(c *gin.Context) {
	param := c.Param("id")
	accountId, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  err,
		})
		return
	}
	collection := models.AccountModel()
	user, err := collection.DeleteOne(context.TODO(), bson.M{"_id": accountId})
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": "success",
		"msg":  "Xóa account thành công",
		"user": user,
	})
}

func ChangeAccountController(c *gin.Context) {
	param := c.Param("id")
	CreatedBy, _ := c.Get("ID")
	accountId, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  err,
		})
		return
	}
	var User InterfaceAccountChangeController
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Data không nhận được",
		})
		return
	}
	User.CreatedBy = CreatedBy
	fmt.Print(User)
	collection := models.AccountModel()
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": accountId}, bson.M{"$set": User}); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "success",
		"msg":  "Thay doi thanh cong",
	})
}
