package config

import (
	"log"

	"github.com/yano-kentaro/todo_app/utils"
	"gopkg.in/go-ini/ini.v1"
)

func init() {
	LoadConfig()
	utils.LogSettings(Config.LogFile)
}

type ConfigList struct {
	Port      string
	SQLDriver string
	DBName    string
	LogFile   string
}

var Config ConfigList

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
	}
}
