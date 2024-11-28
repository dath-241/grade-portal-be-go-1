package controller_admin

import (
	"time"
)

type AuthController struct {
	IDToken string `json:"idToken"`
}

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
	ClassID string `json:"class_id"`
}

type InterfaceAdminController struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Faculty string `json:"faculty"`
	Ms      string `json:"ms"`
}

type InterfaceClassController struct {
	Semester      string   `json:"semester"`
	Name          string   `json:"name"` // nhom lop
	CourseId      string   `json:"course_id"`
	ListStudentMs []string `json:"listStudent_ms"`
	TeacherId     string   `json:"teacher_id"`
	UpdatedBy     any      `json:"updatedBy" bson:"updatedBy"`
}
type InterfaceChangeClassController struct {
	Semester  string `json:"semester"`
	Name      string `json:"name"` // nhom lop
	CourseId  any    `json:"course_id" bson:"course_id"`
	TeacherId any    `json:"teacher_id" bson:"teacher_id"`
	UpdatedBy any    `json:"updatedBy" bson:"updatedBy"`
}

type InterfaceAccountController struct {
	Email     string    `json:"email" bson:"email"`
	Name      string    `json:"name" bson:"name"`
	Ms        string    `json:"ms" bson:"ms"`
	Faculty   string    `json:"faculty" bson:"faculty"`
	Role      string    `json:"role" bson:"role"`
	CreatedBy any       `json:"createdBy" bson:"createdBy"`
	ExpiredAt time.Time `json:"expiredAt" bson:"expiredAt"`
}

type InterfaceCourseController struct {
	Ms     string `json:"ms"`
	Credit int    `json:"credit"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	BT     int    `json:"bt"`
	TN     int    `json:"tn"`
	BTL    int    `json:"btl"`
	GK     int    `json:"gk"`
	CK     int    `json:"ck"`
}

type InterfaceAccountChangeController struct {
	Name      string `json:"name" bson:"name"`
	Faculty   string `json:"faculty" bson:"faculty"`
	Role      string `json:"role" bson:"role"`
	CreatedBy any    `json:"createdBy" bson:"createdBy"`
}

type InterfaceAddStudentClassController struct {
	ClassId       string   `json:"class_id"`
	ListStudentMs []string `json:"listStudent_ms"`
}

// hall of fame
type InterfaceHallOfFame struct {
	Semester string          `json:"semester"`
	Tier     []InterfaceTier `json:"tier"`
}

type InterfaceTier struct {
	CourseID any                    `json:"course_id" bson:"course_id"`
	Data     []InterfaceStudentData `json:"data"`
}

type InterfaceStudentData struct {
	MSSV string  `json:"mssv"`
	DTB  float32 `json:"dtb"`
}

type avgStudentScore struct {
	MSSV     string  `bson:"mssv"`     // Mã số sinh viên
	AvgScore float32 `bson:"avgscore"` // Điểm trung bình
}
