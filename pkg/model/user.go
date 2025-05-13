package model

// リクエスト(/login )時に受け取るJSONデータ
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}