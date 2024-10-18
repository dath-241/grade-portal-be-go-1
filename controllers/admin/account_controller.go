package controller_admin

import (
	"LearnGo/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AccountCreateController(c *gin.Context) {
	var user []InterfaceUserController

	// Bind JSON từ body của request vào struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	createdBy, _ := c.Get("ID")
	// Gán giá trị cho các trường còn thiếu
	for i := range user {
		user[i].CreatedBy = createdBy
		user[i].ExpiredAt = time.Now().AddDate(5, 0, 0) // Ví dụ: hết hạn sau 5 năm
	}

	userCollection := models.UserModel()

	userInterfaces := make([]interface{}, len(user))
	for i := range user {
		userInterfaces[i] = user[i]
	}

	_, err := userCollection.InsertMany(context.TODO(), userInterfaces)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating accounts"})
		return
	}

	// Trả về phản hồi thành công
	c.JSON(http.StatusOK, gin.H{
		"message": "Account created successfully",
		"user":    user,
	})
}
