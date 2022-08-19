package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// データベース接続用のモデルを返す
func NewModels(db *sql.DB) DBModel {
	return DBModel{
		DB: db,
	}
}

type User struct {
	ID        int       `json:"id"`
	LineID	  int    	`json:"line_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type WeightRecord struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	WeightNum int    	`json:"weight_num"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *DBModel) GetOneUser(line_id int) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var u User

	query := `
		select
			id, line_name, created_at, updated_at
		from
			users
		where line_id = ?`

	row := m.DB.QueryRowContext(ctx, query, line_id)

	err := row.Scan(
		&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}

// ユーザの追加
func (m *DBModel) AddUser(line_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		insert into users (line_id, created_at, updated_at)
		values (?, ?, ?)`

	_, err := m.DB.ExecContext(ctx, stmt,
		line_id,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}



// ユーザ情報の更新
func (m *DBModel) UpdateUser(u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		insert into users (first_name, last_name, email, password, created_at, updated_at)
		values (?, ?, ?, ?, ?, ?)`

	_, err := m.DB.ExecContext(ctx, stmt,
		u.LineID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetMaxWeight(id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var maxData int

	query := `
		select
			max(w.weight_num)
		from
			users u
			left join weight_historys w on (u.id = w.user_id)
		where u.id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&maxData,
	)
	if err != nil {
		return 0, err
	}

	return maxData, nil
}
