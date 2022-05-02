//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ config.go [Ver.2022_05_01] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ 開発開始日：2022_05_01
//┃
//┃──┨ 設定管理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                    依存関係
//==========================================|2022_05_01
package config

import (
	"log"

	"github.com/yano-kentaro/todo_app/utils"
	"gopkg.in/go-ini/ini.v1"
)

//===================================================|0
//                    構造体定義
//==========================================|2022_05_01
type ConfigList struct {
	Port      string
	SQLDriver string
	DBName    string
	LogFile   string
	Static    string
}

//===================================================|0
//                    初期化関数
//==========================================|2022_05_01
func init() {
	LoadConfig()
	utils.LogSettings(Config.LogFile)
}

//===================================================|0
//                    グローバル変数
//==========================================|2022_05_01
var Config ConfigList

//===================================================|0
//                    関数定義
//==========================================|2022_05_01
func LoadConfig() {
	conf, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      conf.Section("web").Key("port").MustString("8080"),
		DBName:    conf.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: conf.Section("db").Key("driver").String(),
		LogFile:   conf.Section("web").Key("logfile").String(),
		Static:    conf.Section("web").Key("static").String(),
	}
}
