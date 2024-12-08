package controller_admin

import (
	"LearnGo/models"
	"context"
	"net/http"

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
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "teacher_id không hợp lệ",
		})
		return
	}
	course_id, err := bson.ObjectIDFromHex(data.CourseId)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "course_id không hợp lệ",
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
		"listStudent_ms": data.ListStudentMs,
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
		"code": "success",
		"msg":  "Tạo lớp học thành công",
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

// Hỗ trợ check student hay teacher

func CheckStudentOrTeacher(c *gin.Context, id string, mssv *string) bool { // Student -> true, Teacher -> false
	collection := models.AccountModel()
	// Chuyển đổi id từ string sang ObjectID
	objectId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false // Xử lý lỗi và trả về false
	}

	cursor, err := collection.Find(context.TODO(), bson.M{
		"_id":  objectId,
		"role": "student",
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false // Xử lý lỗi và trả về false
	}
	defer cursor.Close(context.TODO()) // Đảm bảo đóng cursor sau khi sử dụng

	// Kiểm tra xem có tài liệu nào không
	if cursor.Next(context.TODO()) {
		// Nếu có tài liệu, trả về true
		var user models.InterfaceAccount
		cursor.Decode(&user)
		*mssv = user.Ms
		return true
	} else if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false
	}

	// Nếu không có tài liệu nào, trả về false
	return false
}

// API lấy tất cả lớp học theo account_id
func GetAllClassesByAccountID(c *gin.Context) {
	accountID := c.Param("id")

	var classes []bson.M
	collection := models.ClassModel()
	var mssv string

	// Tìm tất cả lớp học mà giáo viên hoặc sinh viên với account_id tham gia
	isStudent := CheckStudentOrTeacher(c, accountID, &mssv)
	var filter bson.M
	if isStudent {
		filter = bson.M{"listStudent_ms": bson.M{"$in": []string{mssv}}} // Nếu là student
	} else {
		id, _ := bson.ObjectIDFromHex(accountID)
		filter = bson.M{"teacher_id": id} // Nếu là teacher
	}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Đọc dữ liệu từ cursor
	for cursor.Next(context.TODO()) {
		var class bson.M
		if err := cursor.Decode(&class); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		classes = append(classes, class)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về danh sách lớp học
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Lấy lớp học thành công",
		"data": gin.H{
			"classes": classes,
		},
	})
}

func GetClassByClassID(c *gin.Context) {
	param := c.Param("id")
	class_id, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid ID format",
		})
		return
	}

	var class models.InterfaceClass
	collection := models.ClassModel()

	if err := collection.FindOne(context.TODO(), bson.M{"_id": class_id}).Decode(&class); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"status": "error",
				"msg":    "Không tìm thấy lớp",
			})
		} else {
			c.JSON(401, gin.H{
				"code": "error",
				"msg":  "ban khong co quyen vao day",
			})
			return
		}
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Lấy lớp học thành công",
		"class":   class,
	})
}

// API lấy tất cả lớp học theo mã môn học
func GetClassByCourseID(c *gin.Context) {
	param := c.Param("id")
	course_id, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid ID format",
		})
		return
	}
	var classes []models.InterfaceClass
	collection := models.ClassModel()
	cursor, err := collection.Find(context.TODO(), bson.M{"course_id": course_id})
	if err != nil {
		c.JSON(400, gin.H{
			"status":  err,
			"message": "Không tìm thấy lớp học",
		})
		return
	}
	for cursor.Next(context.Background()) {
		var class models.InterfaceClass
		if err := cursor.Decode(&class); err != nil {
			c.JSON(400, gin.H{
				"status":  err,
				"message": "Lỗi khi decode dữ liệu",
			})
			return
		}
		classes = append(classes, class)
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Lấy lớp học thành công",
		"classes": classes,
	})
}

func AddStudentsToCourseHandler(c *gin.Context) {
	var request InterfaceAddStudentClassController

	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Data không nhận được",
		})
		return
	}
	collection := models.ClassModel()
	class_id, err := bson.ObjectIDFromHex(request.ClassId)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  err,
		})
	}
	filter := bson.M{"_id": class_id}
	update := bson.M{
		"$addToSet": bson.M{
			"listStudent_ms": bson.M{
				"$each": request.ListStudentMs,
			},
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "error",
			"message": "Failed to add students to course",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "success",
		"message": "Students added to course successfully",
	})
}

func DeleteClassController(c *gin.Context) {
	param := c.Param("id")
	classId, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"code": err,
		})
		return
	}
	collection := models.ClassModel()
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": classId})
	if err != nil {
		c.JSON(500, gin.H{"code": err})
		return
	}
	c.JSON(200, gin.H{
		"code": "success",
		"msg":  "xoa class thanh cong",
	})
}

func ChangeClassController(c *gin.Context) {
	param := c.Param("id")
	classId, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"code": err,
		})
		return
	}
	var data InterfaceChangeClassController
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"code":    "error",
			"message": "Dữ liệu không hợp lệ",
		})
		return
	}
	teacherIdStr, _ := data.TeacherId.(string)
	if teacherIdStr != "" {
		teacher_id, err := bson.ObjectIDFromHex(teacherIdStr)
		if err != nil {
			c.JSON(400, gin.H{
				"code": "error",
				"msg":  "teacher_id không hợp lệ",
			})
			return
		}
		data.TeacherId = teacher_id
	}
	var course_id bson.ObjectID
	courseIdStr, _ := data.CourseId.(string)
	if courseIdStr != "" {
		course_id, err = bson.ObjectIDFromHex(courseIdStr)
		if err != nil {
			c.JSON(400, gin.H{
				"code": "error",
				"msg":  "teacher_id không hợp lệ",
			})
			return
		}
		data.CourseId = course_id
	}
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "course_id không hợp lệ",
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
	data.UpdatedBy = createBy
	class, err := collection.UpdateOne(context.TODO(), bson.M{"_id": classId}, bson.M{"$set": data})

	if err != nil {
		c.JSON(500, gin.H{
			"code":    "error",
			"message": "Lỗi khi tạo lớp học",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":  "success",
		"msg":   "update lop hoc thanh cong",
		"class": class,
	})
}
