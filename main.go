package main

import (
	"log"
	"golang-catch-up/pkg/router" // プロジェクト内のルーター設定パッケージ
	"golang-catch-up/pkg/db"
)

func main() {

	// データベース接続を初期化
  db.InitDB()

	// ルーターのセットアップ
	r := router.SetupRouter()
	/**
		var r *gin.Engine
		r = router.SetupRouter()
	*/

	// サーバーをポート8080とする設定
	port := "8080"

	// 起動時ログにメッセージを出す
	log.Printf("Server started on :%s", port)

	// 指定されたポートでGinサーバーを起動 && エラーチェック
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
/**
	上記の3行は省略記法。きちんと書くと、下記の処理が行われている。

		var err error		// err を外で宣言

	err = r.Run(":"+ port)	// Ginエンジンの Run を呼び出してエラーを取得

		if err != nil {		// err が nil でない場合にエラー処理を行う　→　エラー型に値があるので
		log.Fatalf("Error starting server: %v\n", err)
	}

	「err が nil でない場合って？」  
		・error 型のゼロ値は nil です。
		・なので、err という変数が nil の場合、それは「エラーが発生しなかった」ことを意味する。
		・だから、err に値があるならば（＝errがnillでないならば）処理を行う！
*/

}
