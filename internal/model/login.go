package model

// リクエスト(/login )時に受け取るJSONデータ
// Username: ユーザー名
// Password: パスワード
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
