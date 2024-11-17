package controller_client

import (
	"LearnGo/helper"
	"LearnGo/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetPrevSemesterAllHallOfFame(c *gin.Context) {
	collection := models.HallOfFameModel()
	semester := helper.Set_semester().PREV
	var halloffame_data InterfaceHallOfFame

	var tier_data []InterfaceTier
	filter := bson.M{
		"semester": semester,
	}

	// Sử dụng Find để lấy tất cả các tài liệu khớp với filter
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Đã xảy ra lỗi khi truy vấn dữ liệu"})
		return
	}
	defer cursor.Close(context.TODO())

	// Duyệt qua các tài liệu và thêm chúng vào results
	for cursor.Next(context.TODO()) {
		var data InterfaceTier
		if err := cursor.Decode(&data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Đã xảy ra lỗi khi giải mã dữ liệu"})
			return
		}
		tier_data = append(tier_data, data)
	}

	// Kiểm tra xem có kết quả nào không
	if len(tier_data) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy dữ liệu cho học kỳ trước"})
		return
	} else {
		halloffame_data.Semester = semester
		halloffame_data.Tier = tier_data
	}

	// Trả về tất cả các bản ghi nếu tìm thấy
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Lấy hall of fame thành công",
		"data":    halloffame_data,
	})
}
