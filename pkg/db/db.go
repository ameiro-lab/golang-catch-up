package db

import (
	"context"
	"log"
	"fmt"

	"github.com/jackc/pgx/v4"
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

// GetUsers はデータベースからすべてのユーザー情報を取得する関数
// 
// ユーザーテーブルから id と name のカラムを取得し、マップの配列形式で返却する
// 接続エラーやデータ取得エラーが発生した場合はエラーを返す
//
// 戻り値:
//   - []map[string]interface{}: ユーザー情報を含むマップの配列
//   - error: エラー発生時、詳細なエラー情報を含む
func GetUsers() ([]map[string]interface{}, error) {
	// ユーザーテーブルからIDと名前を取得するSQLクエリ
	query := `SELECT id, user_name FROM "user"`
	
	// データベース接続を使用してクエリを実行
	rows, err := Conn.Query(context.Background(), query)
	if err != nil {
		// クエリ実行エラーの場合、ラップしたエラーを返す
		return nil, fmt.Errorf("クエリ実行に失敗: %w", err)
	}
	// 関数終了時にリソースを確実に解放
	defer rows.Close()
	
	// ユーザー情報を格納するスライスを宣言
	var users []map[string]interface{}
	
	// 結果セットの各行を処理
	for rows.Next() {
		// 各カラムの値を格納する変数を宣言
		var id int
		var name string
		
		// 現在の行のデータを変数にスキャン
		err := rows.Scan(&id, &name)
		if err != nil {
			// スキャンエラーの場合、ラップしたエラーを返す
			return nil, fmt.Errorf("スキャン失敗: %w", err)
		}
		
		// 取得したデータからユーザー情報のマップを作成
		user := map[string]interface{}{
			"id":   id,
			"name": name,
		}
		
		// ユーザー情報をスライスに追加
		users = append(users, user)
	}
	
	// エラーがない場合は結果を返す
	return users, nil
}