package main

import (
	"fmt"
	"log"
	"net/http"	// HTTPサーバーに関するパッケージ
)

/**
	プロコトル＋ホスト名＋ポート(http://localhost:8080)＋"/" にアクセスが来た時、
 */
func handler(w http.ResponseWriter, r *http.Request) {	// 引数の型は importしている net/http パッケージ内で定義された型
	fmt.Fprintf(w, "Hello, World!")	// 第一引数は出力先
}

func main() {

	// リクエスト時の処理を設定する
	http.HandleFunc("/", handler)	// 第一引数に一致するリクエストが来たら、第二引数を実行

	// サーバーをポート8080とする設定
	port := "8080"

	// 起動時ログにメッセージを出す
	fmt.Printf("Server is listening on port %s...\n", port)

	// 指定されたポートでHTTPサーバーを起動 && エラーチェック
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
	/**
		上記の3行は省略記法。きちんと書くと、下記の処理が行われている。

		var err error		// err を外で宣言

		err = http.ListenAndServe(":"+port, nil)	// http.ListenAndServe を呼び出してエラーを取得

		if err != nil {		// err が nil でない場合にエラー処理を行う　→　エラー型に値があるので
				log.Fatalf("Error starting server: %v\n", err)
		}

		「err が nil でない場合って？」
		・error 型のゼロ値は nil です。
		・なので、err という変数が nil の場合、それは「エラーが発生しなかった」ことを意味する。
		・だから、err に値があるならば（＝errがnillでないならば）処理を行う！
	*/

}
