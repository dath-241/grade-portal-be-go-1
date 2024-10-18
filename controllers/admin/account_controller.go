package controller_admin

import (
	"LearnGo/models"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AccountCreateController(c *gin.Context) {
	var newUsers []models.InterfaceUser

	// Bind JSON từ body của request vào struct
	if err := c.ShouldBindJSON(&newUsers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Kết nối với MongoDB
	userCollection := models.UserModel()

	// Lấy tất cả tài khoản từ cơ sở dữ liệu
	var existingUsers []models.InterfaceUser
	cursor, err := userCollection.Find(context.TODO(), map[string][]models.InterfaceUser{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users from database"})
		return
	}
	if err := cursor.All(context.TODO(), &existingUsers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding users"})
		return
	}

	// Các slice để lưu người dùng hợp lệ và bị trùng lặp
	var validUsers []models.InterfaceUser
	var duplicateUsers []models.InterfaceUser
	var invalidEmails []models.InterfaceUser

	// So sánh với các tài khoản hiện có
	for _, newUser := range newUsers {

		if !strings.HasSuffix(newUser.Email, "@hcmut.edu.vn") {
			invalidEmails = append(invalidEmails, newUser)
			continue
		}
		isDuplicate := false
		for _, existingUser := range existingUsers {
			if newUser.Ms == existingUser.Ms || newUser.Email == existingUser.Email {
				// Nếu trùng `ms` hoặc `email`, thêm vào slice duplicateUsers
				duplicateUsers = append(duplicateUsers, newUser)
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			// Thêm thời gian tạo và hết hạn cho các user hợp lệ
			newUser.CreatedAt = time.Now()
			newUser.ExpiredAt = time.Now().AddDate(5, 0, 0) // Ví dụ: hết hạn sau 5 năm
			validUsers = append(validUsers, newUser)
		}
	}

	// Chèn các tài khoản hợp lệ vào cơ sở dữ liệu
	if len(validUsers) > 0 {
		_, err := userCollection.InsertMany(context.TODO(), validUsers)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating valid accounts"})
			return
		}
	}

	// Trả về phản hồi, thông báo người dùng nào đã được thêm và ai bị trùng lặp
	c.JSON(http.StatusOK, gin.H{
		"1.0:_message":                  "Account processing complete",
		"2.0:_valid_users":              validUsers,
		"3.0:_duplicate_users":          duplicateUsers,
		"4.0:_invalid_emails":           invalidEmails,
		"5.0:_total_users":              len(validUsers) + len(duplicateUsers) + len(invalidEmails),
		"6.0:_total_valid":              len(validUsers),
		"7.0:_total_duplicate":          len(duplicateUsers),
		"8.0:_total_invalid_email_form": len(invalidEmails),
	})
}
