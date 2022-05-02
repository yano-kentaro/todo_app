//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ todos.go [Ver.2022_05_01] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ 開発開始日：2022_05_01
//┃
//┃──┨ todosテーブル管理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                    依存関係
//==========================================|2022_05_01
package models

import (
	"log"
	"time"
)

//===================================================|0
//                    構造体定義
//==========================================|2022_05_01
// Todoデータ構造
type Todo struct {
	ID        int
	Content   string
	UserID    string
	CreatedAt time.Time
}

//===================================================|0
//                    Todo新規作成
//==========================================|2022_05_01
func (u *User) CreateTodo(content string) (err error) {
	sql := `
		INSERT INTO todos (
			content, user_id, created_at
		) VALUES (?, ?, ?)
	`
	_, err = DB.Exec(sql, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//===================================================|0
//                    Todo一件取得
//==========================================|2022_05_01
func GetTodo(id int) (todo Todo, err error) {
	sql := `
		SELECT id, content, user_id, created_at
		FROM todos
		WHERE id = ?
	`
	todo = Todo{}
	err = DB.QueryRow(sql, id).Scan(
		&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return todo, err
}

//===================================================|0
//                    Todo全件取得
//==========================================|2022_05_01
func GetTodos() (todos []Todo, err error) {
	sql := `
		SELECT id, content, user_id, created_at
		FROM todos
	`
	rows, err := DB.Query(sql)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	return todos, err
}
