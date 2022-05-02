//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ users.go [Ver.2022_05_01] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ ynkn0829@gmail.com
//┠──┨ 開発開始日：2022_05_01
//┃
//┃──┨ Usersテーブル管理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                 依存関係
//==========================================|2022_05_01
package models

import (
	"log"
	"time"
)

//===================================================|0
//                 構造体定義
//==========================================|2022_05_01
// Userデータ構造
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

//===================================================|0
//                 ユーザー新規作成
//==========================================|2022_05_01
func (u *User) CreateUser() (err error) {
	sql := `
		INSERT INTO users(
			uuid, name, email, password, created_at
		) VALUES (?,?,?,?,?)
	`
	_, err = DB.Exec(sql,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//===================================================|0
//                 ユーザー情報取得
//==========================================|2022_05_01
func GetUser(id int) (user User, err error) {
	user = User{}
	sql := `
		SELECT id, uuid, name, email, password, created_at
		FROM users
		WHERE id = ?
	`
	err = DB.QueryRow(sql, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return user, err
}

//===================================================|0
//                 ユーザー情報編集
//==========================================|2022_05_01
func (u *User) UpdateUser() (err error) {
	sql := `
		UPDATE users SET name = ?, email = ? WHERE id = ?
	`
	_, err = DB.Exec(sql, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//===================================================|0
//                 ユーザー情報削除
//==========================================|2022_05_01
func (u *User) DeleteUser() (err error) {
	sql := `
		DELETE FROM users WHERE id = ?
	`
	_, err = DB.Exec(sql, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
