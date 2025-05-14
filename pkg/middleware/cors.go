package middleware

import "github.com/gin-gonic/gin" // Ginフレームワークのコアパッケージ

// Gin用のCORSミドルウェア（全ルートにCORSヘッダーを追加）
// CORS対応ヘッダーを設定した後、次の処理（muxのルーティング）へリクエストを渡す
// HandlerFunc: *gin.Context を受け取って何も返さない関数の型
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) { // 無名関数を返す => voidなので関数をそのまま返している
		/// CORS用のヘッダーを追加
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // TO DO： "*" ではなく指定ドメインにすべき
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// プリフライトリクエストへの対応
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200) // 200を返して処理を終了
			return
		}

		// 次のミドルウェアやハンドラーを呼び出す(書かなければ次に進まず、後続の処理がスキップされる)
		c.Next()
	}
}
