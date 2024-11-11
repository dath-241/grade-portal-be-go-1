package controller_admin

import (
	"LearnGo/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"LearnGo/helper"
)

func GetPrevSemesterHallOfFame(c *gin.Context) {
	collection := models.HallOfFameModel()
	var data InterfaceHallOfFame
	semester := helper.Set_semester().PREV

	filter := bson.M{
		"semester": semester,
	}
	err := collection.FindOne(context.TODO(), filter).Decode((&data))

	// Kiểm tra lỗi khi tìm kiếm
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy dữ liệu cho học kỳ trước"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Đã xảy ra lỗi khi truy vấn dữ liệu"})
		return
	}

	// Trả về dữ liệu nếu tìm thấy
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Lấy hall of fame thành công",
		"data":    data,
	})
}
