package controller_admin

import (
	"LearnGo/models"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateCourse(c *gin.Context) {
	var data InterfaceCourseController

	// Kiểm tra parse data vào có lỗi không
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"code":    "error",
			"message": "Dữ liệu không hợp lệ",
		})
		return
	}

	collection := models.CourseModel()

	// Kiểm tra xem khóa học có bị trùng không
	isDuplicate, err := CheckDuplicateCourse(collection, data.Ms, data.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "error",
			"message": "Lỗi khi kiểm tra dữ liệu",
		})
		return
	}

	// Nếu khóa học đã tồn tại, trả về lỗi
	if isDuplicate {
		c.JSON(400, gin.H{
			"code":    "error",
			"message": "Khóa học đã tồn tại",
		})
		return
	}

	// Thêm nếu không bị trùng lặp
	createBy, _ := c.Get("ID")
	_, err = collection.InsertOne(context.TODO(), bson.M{
		"ms":        data.Ms,
		"credit":    data.Credit,
		"name":      data.Name,
		"desc":      data.Desc,
		"createdBy": createBy,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"code":    "error",
			"message": "Lỗi khi tạo khóa học",
		})
		return
	}

	// Trả về kết quả thành công
	c.JSON(200, gin.H{
		"code":    "success",
		"message": "Tạo khóa học thành công",
	})
}

func CheckDuplicateCourse(collection *mongo.Collection, ms string, name string) (bool, error) {
	if ms == "" {
		return true, errors.New("Lỗi ms không có")
	}
	filter := bson.M{
		"ms": ms,
	}

	//kiểm tra xem có bản ghi nào không
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil // Không tìm thấy bản ghi
	} else if err != nil {
		return false, err // Có lỗi khác
	}

	return true, nil // Tìm thấy bản ghi trùng
}

// API lấy môn học theo mã môn học
func GetCourseByCourseID(c *gin.Context) {
	param := c.Param("id_course")
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
