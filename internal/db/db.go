package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4" // PostgreSQLドライバ
)

var Conn *pgx.Conn // 外部からアクセスできるように大文字

// 初期化関数
func InitDB() {
	var err error
	Conn, err = pgx.Connect(context.Background(), "postgres://postgres:admin@localhost:5432/microservice")
	if err != nil {
		log.Fatalf("データベース接続に失敗しました: %v", err)
	}
	log.Println("データベース接続成功")
}
