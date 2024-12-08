package controller_admin

import (
	"LearnGo/helper"
	"LearnGo/models"
	"context"
	"errors"
	"fmt"

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

	if data.BT+data.TN+data.BTL+data.GK+data.CK != 100 {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Sai hệ số, tổng hệ số tối đa là 100",
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
		"HS":        [5]int{data.BT, data.TN, data.BTL, data.GK, data.CK},
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
		return true, errors.New("lỗi ms không có")
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

func GetAllCourseController(c *gin.Context) {
	var allCourse []models.InterfaceCourse
	collection := models.CourseModel()
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(400, gin.H{
			"code": err,
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &allCourse); err != nil {
		c.JSON(400, gin.H{
			"code":    "error",
			"message": "Lỗi khi đọc dữ liệu từ cursor",
		})
		return
	}
	semester := helper.Set_semester()
	c.JSON(200, gin.H{
		"code":      "success",
		"msg":       "Lấy ra tất cả khóa học thành công",
		"allCourse": allCourse,
		"semester":  semester,
	})
}

func DeleteCourseController(c *gin.Context) {
	param := c.Param("id")
	course_id, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  err,
		})
	}
	collection := models.CourseModel()
	if _, err = collection.DeleteOne(context.TODO(), bson.M{"_id": course_id}); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  err,
		})
	}
	c.JSON(200, gin.H{
		"code": "success",
		"msg":  "Xóa khóa học thành công",
	})
}

func ChangeCourseController(c *gin.Context) {
	param := c.Param("id")
	course_id, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  1,
		})
		return
	}
	var data struct {
		Ms        string `json:"ms"`
		Credit    int    `json:"credit"`
		Name      string `json:"name"`
		Desc      string `json:"desc"`
		UpdatedBy any    `json:"updatedBy" bson:"updatedBy"`
	}
	if err = c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Data không nhận được",
		})
		return
	}
	collection := models.CourseModel()
	adminId, _ := c.Get("ID")
	data.UpdatedBy = adminId
	fmt.Print(data)
	if _, err = collection.UpdateOne(
		context.TODO(),
		bson.M{
			"_id": course_id,
		},
		bson.M{
			"$set": bson.M{
				"ms":        data.Ms,
				"credit":    data.Credit,
				"name":      data.Name,
				"desc":      data.Desc,
				"updatedBy": data.UpdatedBy,
			},
		}); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  2,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "success",
		"msg":  "Change course thanh cong",
	})
}
