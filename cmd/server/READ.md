├── cmd/
│   └── myapp/
│       └── main.go                 ← エントリーポイント
├── internal/
│   ├── handler/                   ← HTTPハンドラー
│   │   └── auth.go
│   ├── middleware/                ← GinのCORSミドルウェアなど
│   │   └── cors.go
│   ├── repository/                ← DBアクセス
│   │   └── user_repository.go
│   ├── model/                     ← リクエスト・レスポンス構造体など
│   │   └── login.go
│   ├── db/
│   │   └── db.go                  ← DB初期化（Init）
│   └── router/
│       └── router.go              ← Ginエンジン・ルーティング設定
├── go.mod
├── go.sum
