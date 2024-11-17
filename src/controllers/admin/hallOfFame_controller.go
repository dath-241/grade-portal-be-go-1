package controller_admin

import (
	"LearnGo/helper"
	"LearnGo/models"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AvgStudentScores(semester string, course_id bson.ObjectID) ([]avgStudentScore, error) {

	coursesCollection := models.CourseModel()

	// Tìm khóa học với course_id
	var course models.InterfaceCourse
	err := coursesCollection.FindOne(context.TODO(), bson.M{"_id": course_id}).Decode(&course)
	if err != nil {
		return make([]avgStudentScore, 0), errors.New("day la mot loi")
	}
	HS := course.HS
	// Tìm danh sách điểm của sinh viên trong học kỳ và khóa học cụ thể
	scoresCollection := models.ResultScoreModel()
	cursor, err := scoresCollection.Find(context.TODO(), bson.M{"course_id": course_id, "semester": semester})
	if err != nil {
		return make([]avgStudentScore, 0), errors.New("day la mot loi")
	}
	defer cursor.Close(context.TODO())

	var resulScores []models.InterfaceResultScore
	if err = cursor.All(context.TODO(), &resulScores); err != nil {
		return make([]avgStudentScore, 0), errors.New("day la mot loi")
	}
	// Khởi tạo và gán giá trị cho slice avgScores trên cùng một dòng
	totalSize := 0
	for _, result := range resulScores {
		totalSize += len(result.SCORE)
	}
	i := 0
	avgScores := make([]avgStudentScore, totalSize)
	for _, result := range resulScores {
		for _, score := range result.SCORE {
			avgScores[i].MSSV = score.MSSV
			avgScores[i].AvgScore = helper.AvgScore(score.Data, HS[:])
			i++
		}
	}
	return avgScores, nil
}

func MergeSort(avgScores []avgStudentScore) []avgStudentScore {
	if len(avgScores) <= 1 {
		return avgScores
	}

	mid := len(avgScores) / 2
	left := MergeSort(avgScores[:mid])
	right := MergeSort(avgScores[mid:])

	return merge(left, right)
}

func merge(left, right []avgStudentScore) []avgStudentScore {
	var result []avgStudentScore
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].AvgScore > right[j].AvgScore { // Sắp xếp giảm dần
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Thêm các phần tử còn lại
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func CheckDuplicateHOF(collection *mongo.Collection, semester string, course_id bson.ObjectID) bool {

	filter := bson.M{
		"semester":  semester,
		"course_id": course_id,
	}

	//kiểm tra xem có bản ghi nào không
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false // Không tìm thấy bản ghi
	} else if err != nil {
		return false // Có lỗi khác
	}

	return true
}

func CreateHallOfFame(c *gin.Context) {
	scoresCollection := models.ResultScoreModel()
	var results []models.InterfaceResultScore
	semester := helper.Set_semester()
	cursor, err := scoresCollection.Find(context.TODO(), bson.M{
		"semester": semester.PREV,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Loi tim kiem bang ghi",
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &results); err != nil {
		c.JSON(400, gin.H{
			"code": "error",
			"msg":  "Loi tim kiem bang ghi",
		})
		return
	}
	processed := make(map[string]bool)
	collection := models.HallOfFameModel()
	for _, result := range results {
		key := result.Semester + "-" + result.CourseID.Hex()
		if found := processed[key]; !found {
			processed[key] = true
			avgStudentScores, err := AvgStudentScores(result.Semester, result.CourseID)
			if err != nil {
				c.JSON(400, gin.H{
					"code": "error",
					"msg":  err,
				})
				return
			}
			studentHOF := MergeSort(avgStudentScores)
			var data bson.A
			length := min(10, len(studentHOF))
			for i := 0; i < length; i++ {
				student := studentHOF[i]
				data = append(data, bson.M{"mssv": student.MSSV, "dtb": student.AvgScore})
			}
			if !CheckDuplicateHOF(collection, result.Semester, result.CourseID) {

				collection.InsertOne(context.TODO(), bson.M{
					"semester":  result.Semester,
					"course_id": result.CourseID,
					"data":      data,
				})
			} else {
				filter := bson.M{
					"semester":  result.Semester,
					"course_id": result.CourseID,
				}
				update := bson.M{
					"$set": bson.M{
						"data": data},
				}
				collection.UpdateOne(context.TODO(), filter, update)

			}
		}
	}
	c.JSON(200, gin.H{
		"code":    "success",
		"message": "Cập nhật Hall Of Fame thành công",
	})
}

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
