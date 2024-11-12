package controller_client

import (
	"LearnGo/models"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ClassTeacherController(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	if user.Role != "teacher" {
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "Bạn không được quyền vào đây",
		})
		return
	}
	var classTeacherAll []models.InterfaceClass
	collection := models.ClassModel()
	cursor, err := collection.Find(context.TODO(), bson.M{
		"teacher_id": user.ID,
	})
	if err != nil {
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "",
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &classTeacherAll); err != nil {
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "Bạn không được quyền vào đây",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":     "success",
		"classAll": classTeacherAll,
	})
}

func ClassStudentController(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	var classStudentAll []models.InterfaceClassStudent
	collection := models.ClassModel()
	fmt.Println(user)
	cursor, err := collection.Find(context.TODO(), bson.M{
		"listStudent_ms": user.Ms,
	})
	if err != nil {
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "",
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &classStudentAll); err != nil {
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":     "success",
		"classAll": classStudentAll,
	})
}

func ClassAccountController(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	if user.Role == "teacher" {
		ClassTeacherController(c)
		return
	} else if user.Role == "student" {
		ClassStudentController(c)
		return
	}
	c.JSON(400, gin.H{
		"code":    "error",
		"massage": "role nguời dùng không hợp lệ",
	})
}

func ClassDetailController(c *gin.Context) {
	paramID := c.Param("id")
	id, _ := bson.ObjectIDFromHex(paramID)
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	var classDetail models.InterfaceClass
	collection := models.ClassModel()
	err := collection.FindOne(context.TODO(), bson.M{
		"_id": id,
	}).Decode(&classDetail)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    "error",
			"massage": "Không lấy được dữ liệu",
		})
		return
	}
	if user.Role == "student" {
		var listStudent = classDetail.ListStudentMs
		for _, studentMs := range listStudent {
			if studentMs == user.Ms {
				c.JSON(200, gin.H{
					"code":        "success",
					"classDetail": classDetail,
				})
				return
			}
		}
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "Bạn không được phép vào trang này 1",
		})
		return
	} else if user.Role == "teacher" {
		if user.ID != classDetail.TeacherId {
			c.JSON(401, gin.H{
				"code":    "error",
				"massage": "Bạn không được phép vào đây",
			})
			return
		}
		c.JSON(200, gin.H{
			"code":        "success",
			"classDetail": classDetail,
		})
		return
	}
	c.JSON(401, gin.H{
		"code":    "error",
		"massage": "Bạn khồn được phép vào trang này",
	})
}

func CountDocumentController(c *gin.Context) {
	param := c.Param("id")
	courseId, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Không tìm thấy môn học",
		})
		return
	}
	collection := models.ClassModel()
	count, err := collection.CountDocuments(context.TODO(), bson.M{"course_id": courseId})
	if err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Không đếm được các class",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  "success",
		"count": count,
	})
}
