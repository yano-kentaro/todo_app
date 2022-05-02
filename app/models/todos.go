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
	rows.Close()
	return todos, err
}

//===================================================|0
//                    特定ユーザーのTodo取得
//==========================================|2022_05_01
func (u *User) GetTodosByUser(id int) (todos []Todo, err error) {
	sql := `
		SELECT id, content, user_id, created_at
		FROM todos
		WHERE user_id = ?
	`
	rows, err := DB.Query(sql, u.ID)
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
	rows.Close()
	return todos, err
}

//===================================================|0
//                    Todo情報編集
//==========================================|2022_05_01
func (t *Todo) UpdateTodo() (err error) {
	sql := `
		UPDATE todos SET content = ?, user_id = ?
		WHERE id = ?
	`
	_, err = DB.Exec(sql, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//===================================================|0
//                    Todo情報削除
//==========================================|2022_05_01
func (t *Todo) DeleteTodo() (err error) {
	sql := `
		DELETE FROM todos WHERE id = ?
	`
	_, err = DB.Exec(sql, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
