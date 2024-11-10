package controller_client

import (
	"LearnGo/models"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// API lấy môn học theo mã môn học
func GetCourseByCourseID(c *gin.Context) {
	param := c.Param("id")
	course_id, er := bson.ObjectIDFromHex(param)
	if er != nil {
		c.JSON(400, gin.H{
			"code":    "error",
			"message": "ID không hợp lệ",
		})
		return
	}

	var course models.InterfaceCourse
	collection := models.CourseModel()

	if err := collection.FindOne(context.TODO(), bson.M{"_id": course_id}).Decode(&course); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"status":  "error",
				"message": "Không tìm thấy môn học",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Lấy môn học thành công",
		"course":  course,
	})
}
