package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Protocol string
	AppMode string
	HttpPort string
	DB string
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
)

func init() {
	file, err :=  ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Load settings error:", err)
		return
	}
	loadServerSettings(file)
	loadDatabaseSettings(file)
}
func loadServerSettings(file *ini.File)  {
	Protocol = file.Section("server").Key("protocol").MustString("http")
	AppMode = file.Section("server").Key("app_mode").MustString("debug")
	HttpPort = file.Section("server").Key("http_port").MustString("7777")
}
func loadDatabaseSettings(file *ini.File)  {
	DB = file.Section("database").Key("db").MustString("mysql")
	DBHost = file.Section("database").Key("db_host").MustString("localhost")
	DBPort = file.Section("database").Key("db_port").MustString("3306")
	DBUser = file.Section("database").Key("db_user").MustString("tu")
	DBPassword = file.Section("database").Key("db_password").MustString("123456")
	DBName = file.Section("database").Key("db_name").MustString("go_news")
}
