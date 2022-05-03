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
	// 静的ファイルの読み込み
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// ルーティング設定
	http.HandleFunc("/", top)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
