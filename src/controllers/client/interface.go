package controller_client

type InterfaceAccountController struct {
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
type OtpRequest struct {
	Ms string `json:"ms"`
}
