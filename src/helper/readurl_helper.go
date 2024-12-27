package helper

import (
	"LearnGo/models"
	"context"
	"crypto/md5"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type InterfaceScoreController struct {
	BT  []float32 `json:"BT"`
	TN  []float32 `json:"TN"`
	BTL []float32 `json:"BTL"`
	GK  float32   `json:"GK"`
	CK  float32   `json:"CK"`
}

type InterfaceResultScoreController struct {
	SCORE []struct {
		MSSV string                   `json:"mssv"`
		Data InterfaceScoreController `json:"data"`
	} `json:"score"`
	Hash string `json:"hash"`
}

type UploadResultInterface struct {
	ID      bson.ObjectID `json:"_id" bson:"_id"`
	ClassID bson.ObjectID `json:"class_id" bson:"class_id"`
	LinkURL string        `json:"link_url" bson:"file"`
	Hash    string        `json:"hash" bson:"dataHash"`
}

func LoopAuto() {
	for {
		AutoUpdateScore()
		fmt.Println("Đã cập nhật điểm")
		time.Sleep(10 * time.Minute)
	}
}

func AutoUpdateScore() {
	collection := models.ResultScoreModel()
	var AutoUpload []UploadResultInterface
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Lỗi khi lấy dữ liệu từ database")
		return
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &AutoUpload); err != nil {
		fmt.Println("Lỗi khi lấy dữ liệu từ database")
		return
	}
	for _, item := range AutoUpload {
		if item.LinkURL == "" {
			fmt.Println("Không có link URL")
			continue
		}
		scoreDataStr, err := ScoreHelper(item.LinkURL, item.ID.Hex()+".csv", item.Hash)
		fmt.Println(scoreDataStr)
		if err != nil {
			fmt.Println("Lỗi khi lấy dữ liệu từ file CSV")
			continue
		}
		if scoreDataStr == "not change" {
			fmt.Println("Không có thay đổi")
			continue
		}
		var scoreData InterfaceResultScoreController
		err = json.Unmarshal([]byte(scoreDataStr), &scoreData)
		if err != nil {
			fmt.Println("Lỗi khi giải mã JSON")
			continue
		}
		if _, err = collection.UpdateOne(context.TODO(), bson.M{
			"class_id": item.ClassID,
		}, bson.M{
			"$set": bson.M{
				"score":     scoreData.SCORE,
				"expiredAt": time.Now().AddDate(0, 6, 0),
				"dataHash":  scoreData.Hash,
			},
		}); err != nil {
			fmt.Println("Lỗi khi cập nhật dữ liệu vào database")
			continue
		}
	}
}

func ScoreHelper(url string, fileName string, cacheFile string) (string, error) {
	changed, err := checkFileChanged(url, &cacheFile)
	if err != nil {
		fmt.Printf("Lỗi khi kiểm tra file: %v\n", err)
		return "", err
	}

	// Nếu file thay đổi, tải lại và đọc
	if changed {
		fmt.Println("File đã thay đổi, tải lại...")
		err = downloadFile(url, fileName)
		if err != nil {
			fmt.Printf("Lỗi khi tải file: %v\n", err)
		} else {
			fmt.Println("Tải file thành công:", fileName)
		}

		// Làm sạch dữ liệu trong file CSV
		err = cleanCSV(fileName)
		if err != nil {
			fmt.Printf("Lỗi khi làm sạch file CSV: %v\n", err)
		}

		// Đọc file CSV sau khi tải lại
		err = readCSV(fileName)
		if err != nil {
			fmt.Printf("Lỗi khi đọc file CSV: %v\n", err)
		}

		// Đọc file CSV và chuyển đổi dữ liệu
		result, err := readCSVAndConvert(fileName)
		if err != nil {
			fmt.Printf("Lỗi khi đọc file CSV: %v\n", err)
			return "", err
		}
		result.Hash = cacheFile
		// In ra kết quả dưới dạng JSON
		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Printf("Lỗi khi chuyển đổi dữ liệu sang JSON: %v\n", err)
			return "", err
		}
		return string(jsonData), nil
	}
	return "not change", nil
}

func readCSVAndConvert(fileName string) (InterfaceResultScoreController, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return InterfaceResultScoreController{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Cho phép số lượng trường không cố định

	records, err := reader.ReadAll()
	if err != nil {
		return InterfaceResultScoreController{}, fmt.Errorf("lỗi khi đọc file CSV: %v", err)
	}

	var result InterfaceResultScoreController

	// Lấy tiêu đề từ dòng đầu tiên (header)
	header := records[0]

	// Tạo map ánh xạ tiêu đề cột với chỉ số
	headerMap := make(map[string]int)
	for i, title := range header {
		headerMap[title] = i
	}

	// Bỏ qua dòng tiêu đề và đọc các dòng dữ liệu
	for _, record := range records[1:] {
		// Lấy giá trị theo tên tiêu đề
		mssv := record[headerMap["MSSV"]]                          // Ví dụ lấy giá trị của cột MSSV
		bt := parseFloatArrayFromString(record[headerMap["BT"]])   // Ví dụ lấy Score1 -> Score2
		tn := parseFloatArrayFromString(record[headerMap["TN"]])   // Score2 -> Score3
		btl := parseFloatArrayFromString(record[headerMap["BTL"]]) // Score3 -> FinalScore
		gk := parseFloat(record[headerMap["GK"]])
		ck := parseFloat(record[headerMap["CK"]])

		// Tạo đối tượng để lưu dữ liệu
		score := struct {
			MSSV string                   `json:"mssv"`
			Data InterfaceScoreController `json:"data"`
		}{
			MSSV: mssv,
			Data: InterfaceScoreController{
				BT:  bt,
				TN:  tn,
				BTL: btl,
				GK:  gk,
				CK:  ck,
			},
		}

		// Thêm vào kết quả
		result.SCORE = append(result.SCORE, score)
	}

	return result, nil
}
func roundUpToTwoDecimalPlaces(value float64) float32 {
	multiplier := float64(100) // 10^2 để lấy 2 chữ số
	rounded := math.Ceil(float64(value)*multiplier) / multiplier
	return float32(rounded)
}

// Chuyển đổi chuỗi có dạng "8; 8.2" thành mảng float64
func parseFloatArrayFromString(data string) []float32 {
	var result []float32
	values := strings.Split(data, ";") // Tách chuỗi theo dấu ;
	for _, val := range values {
		trimmedVal := strings.TrimSpace(val) // Loại bỏ khoảng trắng
		floatVal, err := strconv.ParseFloat(trimmedVal, 32)
		if err != nil {
			result = append(result, 0) // Nếu không thể chuyển đổi, gán giá trị mặc định là 0
		} else {
			result = append(result, roundUpToTwoDecimalPlaces(floatVal))
		}
	}
	return result
}

// Hàm chuyển đổi một chuỗi thành float32
func parseFloat(field string) float32 {
	value, err := strconv.ParseFloat(strings.TrimSpace(field), 32)
	if err != nil {
		return 0
	}
	return float32(value)
}

// Hàm làm sạch dữ liệu trong file CSV
func cleanCSV(fileName string) error {
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		lines[i] = cleanLine(line)
	}

	output := strings.Join(lines, "\n")
	return os.WriteFile(fileName, []byte(output), 0644)
}

// Hàm làm sạch một dòng trong file CSV
func cleanLine(line string) string {
	// Loại bỏ các dấu ngoặc kép không hợp lệ
	return strings.ReplaceAll(line, `"`, ``)
}

// Hàm đọc file CSV
func readCSV(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		_, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("lỗi khi đọc file CSV: %v", err)
		}
	}

	return nil
}

// Hàm kiểm tra file có thay đổi không (dựa vào hash)
func checkFileChanged(url string, cacheFile *string) (bool, error) {
	// Tải file từ URL
	resp, err := http.Get(url)
	if err != nil {
		return false, fmt.Errorf("lỗi khi tải file từ URL: %v", err)
	}
	defer resp.Body.Close()

	// Tạo file tạm thời để lưu nội dung tải về
	tempFile, err := os.CreateTemp("", "tempFile")
	if err != nil {
		return false, fmt.Errorf("lỗi khi tạo file tạm thời: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Ghi nội dung tải về vào file tạm thời
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return false, fmt.Errorf("lỗi khi ghi nội dung vào file tạm thời: %v", err)
	}

	// Tính hash của file tạm thời
	tempFileHash, err := calculateFileHash(tempFile.Name())
	if err != nil {
		return false, fmt.Errorf("lỗi khi tính hash của file tạm thời: %v", err)
	}

	if tempFileHash != *cacheFile {
		*cacheFile = tempFileHash
		return true, nil
	}

	return false, nil
}

// Hàm tính hash của file
func calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// Hàm tải file từ URL
func downloadFile(url, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
