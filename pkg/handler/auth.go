package handler

import (
	"fmt"                       // 書式付き入出力のための標準パッケージ（デバッグ出力に使用）
	"golang-catch-up/pkg/model" // データモデル定義を含むプロジェクト内パッケージ
	"net/http"                  // HTTPステータスコードを使用するため

	"github.com/gin-gonic/gin" // Ginフレームワークのコアパッケージ
	"golang-catch-up/pkg/db"
)

// "/login"のPOSTリクエストを処理する
// *gin.Context: リクエスト情報・レスポンス用の情報
func LoginHandler(c *gin.Context) {

	// LoginRequest型をゼロ値で初期化
	var loginData model.LoginRequest
	/**
		loginData := model.LoginRequest{}　と書いても
		loginData := model.LoginRequest{
	    Username: "test",
	    Password: "pass",
		}	と初期値を含めてもOK
	*/

	// JSONデータのバインド && エラーチェック
	if err := c.ShouldBindJSON(&loginData); err != nil {
		// 失敗した場合はエラーレスポンスを返す
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 受信したログイン情報をコンソールに出力する
	fmt.Printf("Received login request: %s, %s\n", loginData.Username, loginData.Password)

	// DBアクセス && エラーチェック
	users, err := db.GetUsers()
	if err != nil {
		// エラーログを追加して詳細な情報を出力
		fmt.Printf("Error fetching users from DB: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザー情報の取得に失敗しました"})
		return
	}

	// 成功レスポンスを返す
	c.JSON(http.StatusOK, gin.H{"users": users}) // TO DO：認証処理を実装する
}
