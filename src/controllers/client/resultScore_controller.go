package controller_client

import (
	"LearnGo/models"
	"context"
	"log"

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
	var data InterfaceResultScoreController
	// lay du lieu tu front end
	c.BindJSON(&data)

}
