package controller_admin

import (
	"LearnGo/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func AccountCreateController(c *gin.Context) {
	var newUsers []InterfaceUserController
	// Bind JSON từ body của request vào struct
	if err := c.ShouldBindJSON(&newUsers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	userCollection := models.UserModel()

	// Lấy tất cả tài khoản từ cơ sở dữ liệu
	var existingUsers []models.InterfaceUser
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
	c.JSON(200, gin.H{
		"code":    "success",
		"massage": CreatedBy,
	})
	m := make(map[string]bool)
	var filterAccount []InterfaceUserController
	var errorAccount []InterfaceUserController
	for _, account := range existingUsers {
		m[account.Email] = true
		m[account.Ms] = true
	}
	for _, newAccount := range newUsers {
		if !m[newAccount.Email] && !m[newAccount.Ms] {
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
		"code":         "success",
		"errorAccount": errorAccount,
	})
}
