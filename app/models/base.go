//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ base.go [Ver.2022_05_01] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ ynkn0829@gmail.com
//┠──┨ 開発開始日：2022_05_01
//┃
//┃──┨ DB管理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                    依存関係
//==========================================|2022_05_01
package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yano-kentaro/todo_app/config"
)

//===================================================|0
//                    グローバル変数
//==========================================|2022_05_01
var DB *sql.DB
var err error

const (
	tableNameUsers = "users"
	tableNameTodos = "todos"
)

//===================================================|0
//                    テーブル構造
//==========================================|2022_05_01

// usersテーブル
const tableStructUsers = `
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid STRING NOT NULL UNIQUE,
	name STRING,
	email STRING,
	password STRING,
	created_at DATETIME
`

// todosテーブル
const tableStructTodos = `
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	content TEXT,
	user_id INTEGER,
	created_at DATETIME
`

//===================================================|0
//                    初期化関数
//==========================================|2022_05_01
func init() {
	DB, err = sql.Open(config.Config.SQLDriver, config.Config.DBName)
	if err != nil {
		log.Fatalln(err)
	}

	sqlUser := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(%s)`, tableNameUsers, tableStructUsers)
	DB.Exec(sqlUser)

	sqlTodos := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(%s)`, tableNameTodos, tableStructTodos)
	DB.Exec(sqlTodos)
}

//===================================================|0
//                    UUID作成
//==========================================|2022_05_01
func createUUID() (uuidObj uuid.UUID) {
	uuidObj, _ = uuid.NewUUID()
	return uuidObj
}

//===================================================|0
//                    暗号化関数
//==========================================|2022_05_01
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
