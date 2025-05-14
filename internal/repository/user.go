package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4" // PostgreSQLドライバ
)

type User struct {
	ID   int
	Name string
}

type UserRepository struct {
	conn *pgx.Conn
}

// コンストラクタ
func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{conn: conn}
}

// userテーブルから全件取得する
func (r *UserRepository) GetAll(ctx context.Context) ([]User, error) {
	rows, err := r.conn.Query(ctx, `SELECT id, user_name FROM "user"`)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行に失敗: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, fmt.Errorf("スキャン失敗: %w", err)
		}
		users = append(users, u)
	}

	return users, nil
}
