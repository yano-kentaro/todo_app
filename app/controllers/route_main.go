//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ route_main.go [Ver.2022_05_01] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ 開発開始日：2022_05_01
//┃
//┃──┨ HTMLファイルのルーティング管理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                    依存関係
//==========================================|2022_05_01
package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yano-kentaro/todo_app/app/models"
)

//===================================================|0
//                    "/"へのアクセス
//==========================================|2022_05_01
func top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func returnTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetTodos()
	if err != nil {
		log.Fatalln(err)
	}
	json.NewEncoder(w).Encode(todos)
	fmt.Println("Endpoint: /todos")
}
