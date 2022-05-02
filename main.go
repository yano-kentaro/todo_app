//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ main.go [Ver.2022_05_01] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ ynkn0829@gmail.com
//┠──┨ 開発開始日：2022_05_01
//┃
//┃──┨ メイン処理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//					依存関係
//==========================================|2022_05_01
package main

import (
	"fmt"
	// "log"

	"github.com/yano-kentaro/todo_app/app/models"
	// "github.com/yano-kentaro/todo_app/config"
)

//===================================================|0
//					メイン処理
//==========================================|2022_05_01
func main() {
	fmt.Print("================================出力結果================================\n\n")

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "qwe123"
	// fmt.Println(u)
	// u.CreateUser()

	u, _ := models.GetUser(1)
	fmt.Println(u)

	fmt.Print("\n========================================================================\n")
}
