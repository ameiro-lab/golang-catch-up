package main

import (
	"log"
	"net/http"	// HTTPサーバーに関するパッケージ
	"golang-catch-up/pkg/handler"
)

// CORSヘッダを追加するミドルウェア（HTTPリクエスト時、最初に実行される処理）
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS用のヘッダーを追加
		w.Header().Set("Access-Control-Allow-Origin", "*") // TO DO： "*" ではなく指定ドメインにすべき
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// プリフライトリクエストへの対応
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// CORS対応ヘッダーを設定した後、次の処理（muxのルーティング）にリクエストを渡す
		next.ServeHTTP(w, r)
	})
}

func main() {

	// マルチプレクサ(mux)ルーターの作成
	mux := http.NewServeMux()

	// mux にエンドポイントの登録
	mux.HandleFunc("/", handler.HelloHandler)						// "/" にアクセスされたときに HelloHandler を実行
	mux.HandleFunc("/api/login", handler.LoginHandler)	// "/api/login" にアクセスされたときに LoginHandler を実行

	// mux にCORS対応ミドルウェアを設定する
	handlerWithCORS := corsMiddleware(mux)
	/**
		var handlerWithCORS http.Handler
		handlerWithCORS = corsMiddleware(mux)
	*/

	// サーバーをポート8080とする設定
	port := "8080"

	// 起動時ログにメッセージを出す
	log.Printf("Server started on :%s", port)

	// 指定されたポートでHTTPサーバーを起動 && ルーティング && エラーチェック
	if err := http.ListenAndServe(":" + port, handlerWithCORS); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
	/**
		上記の3行は省略記法。きちんと書くと、下記の処理が行われている。

		var err error		// err を外で宣言

		err = http.ListenAndServe(":"+ port, handlerWithCORS)	// http.ListenAndServe を呼び出してエラーを取得

		if err != nil {		// err が nil でない場合にエラー処理を行う　→　エラー型に値があるので
				log.Fatalf("Error starting server: %v\n", err)
		}

		「err が nil でない場合って？」
		・error 型のゼロ値は nil です。
		・なので、err という変数が nil の場合、それは「エラーが発生しなかった」ことを意味する。
		・だから、err に値があるならば（＝errがnillでないならば）処理を行う！
	*/

}
