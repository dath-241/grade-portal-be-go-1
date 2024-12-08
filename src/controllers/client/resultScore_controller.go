package controller_client

import (
	"LearnGo/models"
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ResultController(c *gin.Context) {
	data, _ := c.Get("user")
	param := c.Param("id")
	class_id, _ := bson.ObjectIDFromHex(param)
	user := data.(models.InterfaceAccount)
	var resultScore models.InterfaceResultScore
	collection := models.ResultScoreModel()
	if err := collection.FindOne(context.TODO(), bson.M{
		"class_id": class_id,
	}).Decode(&resultScore); err != nil {
		c.JSON(401, gin.H{
			"code": "error",
			"msg":  "ban khong co quyen vao day",
		})
		return
	}
	if user.Role == "teacher" {
		c.JSON(200, gin.H{
			"code":        "success",
			"resultScore": resultScore,
		})
		return
	} else if user.Role == "student" {
		for _, item := range resultScore.SCORE {
			if item.MSSV == user.Ms {
				c.JSON(200, gin.H{
					"code":  "success",
					"score": item,
				})
				return
			}

		}
	}
	c.JSON(401, gin.H{
		"code": "error",
		"msg":  "ban khong co quyen vao day",
	})
}

func CreateResultScoreController(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	var dataResult InterfaceResultScoreController
	// lay du lieu tu front end
	if err := c.BindJSON(&dataResult); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Data không nhận được",
		})
		return
	}
	class_id, err := bson.ObjectIDFromHex(dataResult.ClassID)
	if err != nil {
		c.JSON(204, gin.H{
			"code": "error",
			"msg":  "Lớp chưa có giáo viên",
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
	if _, err = collection.InsertOne(context.TODO(), bson.M{
		"semester":  classDetail.Semester,
		"course_id": classDetail.CourseId,
		"score":     dataResult.SCORE,
		"class_id":  class_id,
		"expiredAt": time.Now().AddDate(0, 6, 0),
		"createdBy": user.ID,
		"updatedBy": user.ID,
	}); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Cap nhat bang diem thất bại",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "success",
		"msg":  "Cap nhat bang diem thanh cong",
	})
}

func ResultPatchController(c *gin.Context) {
	id := c.Param("id")
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	var dataResult InterfaceResultScoreController
	if err := c.BindJSON(&dataResult); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Data không nhận được",
		})
		return
	}
	class_id, _ := bson.ObjectIDFromHex(id)
	collection := models.ResultScoreModel()
	result, err := collection.UpdateOne(
		context.TODO(),
		bson.M{
			"class_id": class_id,
		},
		bson.M{
			"$set": bson.M{
				"score":     dataResult.SCORE,
				"updatedBy": user.ID,
			},
		},
	)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "loi",
		})
		return
	}

	if result.MatchedCount != 0 {
		c.JSON(200, gin.H{
			"code": "success",
			"msg":  "Thay đổi thành công",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "success",
		"msg":  "Thay đổi thành công",
	})
}

func ResultCourseController(c *gin.Context) {
	data, _ := c.Get("user")
	account := data.(models.InterfaceAccount)
	param := c.Param("ms")
	params := strings.Split(param, "-")
	var course models.InterfaceCourse
	collection_course := models.CourseModel()
	if err := collection_course.FindOne(context.TODO(), bson.M{"ms": params[0]}).Decode(&course); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "MS course sai",
		})
		return
	}
	var resultScore models.InterfaceResultScore
	collection_result := models.ResultScoreModel()
	if err := collection_result.FindOne(context.TODO(), bson.M{"course_id": course.ID, "semester": params[1]}).Decode(&resultScore); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "ID course sai",
		})
		return
	}
	for _, item := range resultScore.SCORE {
		if item.MSSV == account.Ms {
			c.JSON(200, gin.H{
				"code":  "success",
				"msg":   "Lấy điểm thành công",
				"name":  course.Name,
				"score": item.Data,
			})
			return
		}
	}
	c.JSON(400, gin.H{
		"code": "error",
		"msg":  "",
	})
}

func ResultAllController(c *gin.Context) {
	data, _ := c.Get("user")
	account := data.(models.InterfaceAccount)
	collection := models.ResultScoreModel()
	var result []models.InterfaceResultScore
	cursor, err := collection.Find(context.TODO(), bson.M{
		"score.mssv": account.Ms,
	})
	if err != nil {
		c.JSON(401, gin.H{
			"code": "error",
			"msg":  "3",
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &result); err != nil {
		c.JSON(401, gin.H{
			"code": "error",
			"msg":  "4",
		})
		return
	}
	type score struct {
		Ms   string                `json:"ms"`
		Name string                `json:"name"`
		Data models.InterfaceScore `json:"data"`
	}
	var scores []score
	for _, item := range result {
		for _, sco := range item.SCORE {
			if sco.MSSV == account.Ms {
				collection_course := models.CourseModel()
				var course models.InterfaceCourse
				if err := collection_course.FindOne(context.TODO(), bson.M{"_id": item.CourseID}).Decode(&course); err != nil {
					c.JSON(400, gin.H{
						"code": "error",
						"msg":  "MS course sai",
					})
					return
				}
				scores = append(scores, score{course.MS + "-" + item.Semester, course.Name, sco.Data})
			}
		}
	}
	c.JSON(200, gin.H{
		"code":   "success",
		"msg":    "Lấy điểm thành công",
		"scores": scores,
	})
}
