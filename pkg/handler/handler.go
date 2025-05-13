package handler

import (
	"fmt"
	"net/http"
)

// "/"のリクエストを処理する
func HelloHandler(w http.ResponseWriter, r *http.Request) {	// 引数の型は importしている net/http パッケージ内で定義された型
	fmt.Fprintf(w, "Hello, World!")	// 第一引数は出力先
}
