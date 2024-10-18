package controller_admin

import (
	"LearnGo/models"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateClass(c *gin.Context) {
	var data InterfaceClass

	c.BindJSON(&data)
	collection := models.ClassModel()
	createAt, _ := c.Get("ID")
	collection.InsertOne(context.TODO(), bson.M{
		"semester":       data.Semester,
		"name":           data.Name,
		"course_id":      data.CourseId,
		"listStudent_id": data.ListStudentId,
		"teacher_id":     data.TeacherId,
		"createdBy":      createAt,
	})

	c.JSON(200, gin.H{
		"code":    "success",
		"message": "Tạo lớp học thành công",
	})
}
