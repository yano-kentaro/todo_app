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
	"html/template"
	"log"
	"net/http"
)

//===================================================|0
//                    "/"へのアクセス
//==========================================|2022_05_01
func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.
		New("top.html").
		Delims("[[", "]]"). // Vueのマスタッシュを機能させるため
		ParseFiles("./app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, "test") // 第二引数の値をHTMLへ渡すことが出来る
}
