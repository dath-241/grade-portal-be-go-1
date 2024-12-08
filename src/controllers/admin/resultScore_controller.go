package controller_admin

import (
	"LearnGo/models"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ResultScoreController(c *gin.Context) {
	var data InterfaceResultScoreController
	// lay du lieu tu front end
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Data không nhận được",
		})
		return
	}
	class_id, err := bson.ObjectIDFromHex(data.ClassID)
	if err != nil {
		c.JSON(204, gin.H{
			"code": "error",
			"msg":  "Lớp chưa có giáo viên",
		})
	}
	createBy, _ := c.Get("ID")
	collection := models.ResultScoreModel()
	var ResultScore models.InterfaceResultScore
	err = collection.FindOne(
		context.TODO(),
		bson.M{
			"class_id": class_id,
		},
	).Decode(&ResultScore)
	// co ban ghi resultScore truoc do
	if err == nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Bảng ghi của lớp học này đã được lưu trong database trước đó",
		})
		return
	}

	var classDetail models.InterfaceClass
	collectionClass := models.ClassModel()

	if err = collectionClass.FindOne(context.TODO(), bson.M{"_id": class_id}).Decode(&classDetail); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Khong tim thay lop hoc do",
		})
		return
	}
	collection.InsertOne(context.TODO(), bson.M{
		"semester":  classDetail.Semester,
		"course_id": classDetail.CourseId,
		"score":     data.SCORE,
		"class_id":  class_id,
		"expiredAt": time.Now().AddDate(0, 6, 0),
		"createdBy": createBy,
		"updatedBy": createBy,
	})
	c.JSON(200, gin.H{
		"code": "success",
		"msg":  "Cap nhat bang diem thanh cong",
	})
}

func GetResultScoreController(c *gin.Context) {
	param := c.Param("id")
	class_id, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  err,
		})
		return
	}
	collection := models.ResultScoreModel()
	var data models.InterfaceResultScore
	if err = collection.FindOne(context.TODO(), bson.M{"class_id": class_id}).Decode(&data); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  "success",
		"msg":   "Lấy bảng điểm thành công",
		"score": data,
	})
}
