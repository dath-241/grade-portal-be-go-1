package controller_admin

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
