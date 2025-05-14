package router

import (
	"golang-catch-up/internal/handler"
	"golang-catch-up/internal/middleware"

	"github.com/gin-gonic/gin" // Ginフレームワークのコアパッケージ
)

// ルーターをセットアップする
// Ginエンジン（*gin.Engine 型）を初期化し、結果を返す関数
func SetupRouter() *gin.Engine {

	// Ginエンジンをデフォルト設定で初期化する（ロギングミドルウェア・リカバリーミドルウェア込み）
	r := gin.Default()

	// Ginエンジン全体にCORSミドルウェア(カスタム)を適用する
	r.Use(middleware.CORSMiddleware())

	// ルーティング
	api := r.Group("/api") // API通信用のエンドポイントグループを生成
	{
		api.POST("/login", handler.LoginHandler) // "/api/login" にアクセスされたときに LoginHandler を実行
	}

	// トップページ(ルートパス)にアクセスされた時の挙動を設定
	// r.GET("/", handler.HelloHandler)	// 開発時は不要なのでひとまずコメントアウト

	// 設定済みのGinエンジンを返却
	return r
}
