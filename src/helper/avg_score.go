package helper

import (
	"LearnGo/models"
)

// AvgScore tính toán điểm trung bình tổng hợp từ các điểm BT, TN và BTL
func AvgScore(data models.InterfaceScore, HS []int) float32 {
	var avgscore float32

	// Tính điểm trung bình BT
	var totalBT float32
	var countBT int
	for _, score := range data.BT {
		totalBT += score
		countBT++
	}
	var avgBT float32
	if countBT > 0 {
		avgBT = totalBT / float32(countBT)
	} else {
		avgBT = 0
	}

	// Tính điểm trung bình TN
	var totalTN float32
	var countTN int
	for _, score := range data.TN {
		totalTN += score
		countTN++
	}
	var avgTN float32
	if countTN > 0 {
		avgTN = totalTN / float32(countTN)
	} else {
		avgTN = 0
	}

	// Tính điểm trung bình BTL
	var totalBTL float32
	var countBTL int
	for _, score := range data.BTL {
		totalBTL += score
		countBTL++
	}
	var avgBTL float32
	if countBTL > 0 {
		avgBTL = totalBTL / float32(countBTL)
	} else {
		avgBTL = 0
	}

	// Tính điểm trung bình tổng hợp (avgscore) dựa trên hệ số HS
	avgscore = float32(HS[0])*avgBT/100 + float32(HS[1])*avgTN/100 + float32(HS[2])*avgBTL/100 + float32(HS[3])*data.GK/100 + float32(HS[4])*data.CK/100

	return avgscore
}
