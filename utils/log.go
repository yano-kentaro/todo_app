//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ log.go [Ver.2022_05_01] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ ynkn0829@gmail.com
//┠──┨ 開発開始日：2022_05_01
//┃
//┠──┨ ログ出力設定の管理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//					依存関係
//==========================================|2022_05_01
package utils

import (
	"io"
	"log"
	"os"
)

//===================================================|0
//					ログ設定
//==========================================|2022_05_01
func LogSettings(logFile string) {
	lFile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, lFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
