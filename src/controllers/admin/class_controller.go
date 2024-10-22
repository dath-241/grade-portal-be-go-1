package controller_admin

import (
	"LearnGo/models"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateClass(c *gin.Context) {
	var data InterfaceClassController
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"code":    "error",
			"message": "Dữ liệu không hợp lệ",
		})
		return
	}
	teacher_id, err := bson.ObjectIDFromHex(data.TeacherId)
	fmt.Println(teacher_id, data)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    "error",
			"massage": "teacher_id không hợp lệ",
		})
		return
	}
	course_id, err := bson.ObjectIDFromHex(data.CourseId)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    "error",
			"massage": "course_id không hợp lệ",
		})
		return
	}
	// Kiểm tra parse data vào có lỗi ko

	collection := models.ClassModel()

	// Kiểm tra xem lớp học có bị trùng ko bằng FindOne
	isDuplicate, err := CheckDuplicateClass(collection, data.Semester, course_id, data.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "error",
			"message": "Lỗi khi kiểm tra dữ liệu",
		})
		return
	}

	// Nếu lớp học đã tồn tại, trả về lỗi
	if isDuplicate {
		c.JSON(400, gin.H{
			"code":    "error",
			"message": "Lớp học đã tồn tại",
		})
		return
	}
	// Thêm nếu không bị trùng lăp
	createBy, _ := c.Get("ID")

	_, err = collection.InsertOne(context.TODO(), bson.M{
		"semester":       data.Semester,
		"name":           data.Name,
		"course_id":      course_id,
		"listStudent_id": data.ListStudentId,
		"teacher_id":     teacher_id,
		"createdBy":      createBy,
		"updatedBy":      createBy,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"code":    "error",
			"message": "Lỗi khi tạo lớp học",
		})
		return
	}

	// Trả về kết quả thành công
	c.JSON(200, gin.H{
		"code":    "success",
		"message": "Tạo lớp học thành công",
	})
}

func CheckDuplicateClass(collection *mongo.Collection, semester string, courseId bson.ObjectID, name string) (bool, error) {
	filter := bson.M{
		"semester":  semester,
		"course_id": courseId,
		"name":      name,
	}

	// Sử dụng FindOne để kiểm tra xem có bản ghi nào không
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil // Không tìm thấy bản ghi
	} else if err != nil {
		return false, err // Có lỗi khác
	}

	return true, nil // Tìm thấy bản ghi trùng
}
