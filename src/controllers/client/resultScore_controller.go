package controller_client

import (
	"LearnGo/models"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ResultController(c *gin.Context) {
	data, _ := c.Get("user")
	param := c.Param("id")
	class_id, _ := bson.ObjectIDFromHex(param)
	user := data.(models.InterfaceUser)
	var resultScore models.InterfaceResultScore
	collection := models.ResultScoreModel()
	if err := collection.FindOne(context.TODO(), bson.M{
		"class_id": class_id,
	}).Decode(&resultScore); err != nil {
		log.Fatalf("Find error: %v", err)
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
		"code":    "error",
		"massage": "ban khong co quyen vao day",
	})
}

func CreateResultScoreController(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceUser)
	var dataResult InterfaceResultScoreController
	// lay du lieu tu front end
	c.BindJSON(&dataResult)
	class_id, err := bson.ObjectIDFromHex(dataResult.ClassID)
	if err != nil {
		c.JSON(204, gin.H{
			"code":    "error",
			"massage": "Lớp chưa có giáo viên",
		})
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
			"code":    "error",
			"massage": "Bảng ghi của lớp học này đã được lưu trong database trước đó",
		})
		return
	}
	collection.InsertOne(context.TODO(), bson.M{
		"score":     dataResult.SCORE,
		"class_id":  class_id,
		"expiredAt": time.Now().AddDate(0, 6, 0),
		"createdBy": user.ID,
		"updatedBy": user.ID,
	})
	c.JSON(200, gin.H{
		"code":    "success",
		"massage": "Cap nhat bang diem thanh cong",
	})
}

func ResultPatchController(c *gin.Context) {
	id := c.Param("id")
	data, _ := c.Get("user")
	user := data.(models.InterfaceUser)
	var dataResult InterfaceResultScoreController
	c.BindJSON(&dataResult)
	class_id, _ := bson.ObjectIDFromHex(id)
	collection := models.ResultScoreModel()
	result, err := collection.UpdateOne(
		context.TODO(),
		bson.M{
			"class_id": class_id,
		},
		bson.M{
			"score":     dataResult.SCORE,
			"updatedBy": user.ID,
		},
	)
	if err != nil {
		log.Panic(err)
	}

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return
	}
	c.JSON(200, gin.H{
		"code":    "success",
		"massage": "Thay đổi thành công",
	})

}
