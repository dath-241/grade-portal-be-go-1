package controller_client

import (
	"LearnGo/models"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ClassTeacherController(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceUser)
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
		log.Fatalf("Find error: %v", err)
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &classTeacherAll); err != nil {
		log.Fatalf("Cursor All error: %v", err)
	}

	c.JSON(200, gin.H{
		"code":     "success",
		"classAll": classTeacherAll,
	})
}

func ClassStudentController(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceUser)
	var classStudentAll []models.InterfaceClass
	collection := models.ClassModel()
	fmt.Println(user)
	cursor, err := collection.Find(context.TODO(), bson.M{
		"listStudent_id": user.ID.Hex(),
	})
	if err != nil {
		log.Fatalf("Find error: %v", err)
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &classStudentAll); err != nil {
		log.Fatalf("Cursor All error: %v", err)
	}
	c.JSON(200, gin.H{
		"code":     "success",
		"classAll": classStudentAll,
	})
}

func ClassDetailController(c *gin.Context) {
	paramID := c.Param("id")
	id, _ := bson.ObjectIDFromHex(paramID)
	data, _ := c.Get("user")
	user := data.(models.InterfaceUser)
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
		var listStudent = classDetail.ListStudentId
		for _, studentID := range listStudent {
			if studentID == user.ID.String() {
				c.JSON(200, gin.H{
					"code":        "success",
					"classDetail": classDetail,
				})
				return
			}
		}
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "Bạn không được phép vào trang này",
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
