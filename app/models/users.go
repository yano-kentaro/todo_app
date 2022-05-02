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
//					依存関係
//==========================================|2022_05_01
package models

import (
	"log"
	"time"
)

//===================================================|0
//					構造体定義
//==========================================|2022_05_01
// Usersテーブル
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

//===================================================|0
//					ユーザー新規作成
//==========================================|2022_05_01
func (u *User) CreateUser() (err error) {
	sql := `
				insert into users(
					uuid, name, email, password, created_at
				) values (?,?,?,?,?)
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
//					ユーザー編集
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
