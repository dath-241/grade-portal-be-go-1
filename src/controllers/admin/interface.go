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
		MSSV string                   `json:"MMSV"`
		Data InterfaceScoreController `json:"Data"`
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
	ListStudentId []string `json:"listStudent_id"`
	TeacherId     string   `json:"teacher_id"`
}

type InterfaceUserController struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Ms        string    `json:"ms"`
	Faculty   string    `json:"faculty"`
	Role      string    `json:"role"`
	CreatedBy any       `json:"createdBy" bson:"createdBy"`
	ExpiredAt time.Time `json:"expiredAt" bson:"createdAt"`
}

type InterfaceCourseController struct {
	Ms     string `json:"ms"`
	Credit string `json:"credit"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
}
