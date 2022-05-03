//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ server.go [Ver.2022_05_01] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ 開発開始日：2022_05_01
//┃
//┃──┨ WEBサーバー管理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                    依存関係
//==========================================|2022_05_01
package controllers

import (
	"net/http"

	"github.com/yano-kentaro/todo_app/config"
)

//===================================================|0
//                    メインサーバー起動
//==========================================|2022_05_01
func StartMainServer() (err error) {
	// ルーティング設定
	http.HandleFunc("/", top)
	http.HandleFunc("/todos/", returnTodos)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}

// //===================================================|0
// //                    HTML生成
// //==========================================|2022_05_01
// func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
// 	var files []string
// 	for _, file := range filenames {
// 		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
// 	}
// 	// var templates *template.Template
// 	// for _, t := range files {
// 	// 	templates = append(
// 	// 		templates, template.New(t).Delims("[[", "]]").ParseFiles(t),
// 	// 	)
// 	// }
// 	templates := template.Must(
// 		template.New("layout").Delims("[[", "]]").ParseFiles(files...),
// 	)
// 	templates.ExecuteTemplate(w, "layout", data)
// }
