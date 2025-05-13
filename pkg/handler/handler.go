package handler

import (
	"fmt"
	"net/http"

	"encoding/json"	// JSONとの変換（エンコード・デコード） を扱うための標準パッケージ
	"golang-catch-up/pkg/model" // modelパッケージをインポート
)

// "/"のリクエストを処理する
func HelloHandler(w http.ResponseWriter, r *http.Request) {	// 引数の型は importしている net/http パッケージ内で定義された型
	fmt.Fprintf(w, "Hello, World!")	// 第一引数は出力先
}

// "/login"のPOSTリクエストを処理する
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// POSTメソッド以外のリクエストを拒否する
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	
	// リクエストボディからログイン情報を取得
	var loginData model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		// JSONデコードエラーの場合は400 Bad Requestを返す
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	
	// ログに出力
	fmt.Printf("Received login request: %s, %s\n", loginData.Username, loginData.Password)
	
	// TODO: 実際のユーザー認証処理を実装する
	
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// 成功レスポンスの返却
	response := map[string]string{"message": "Login successful"}
	json.NewEncoder(w).Encode(response)
}