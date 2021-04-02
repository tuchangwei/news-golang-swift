package utils

import (
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"strings"
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

	var env string
	flag.StringVar(&env, "env", "dev", "\"dev\" or \"pro\"")
	flag.Parse()

	env = strings.ToLower(env)
	var file *ini.File
	var err error
	switch env {
	case "dev":
		fmt.Println("You are under Development environment")
		file, err =  ini.Load("config/config.ini")
	case "pro":
		fmt.Println("You are under Production environment")
		file, err =  ini.Load("../../news-golang-swift-production-config/config.ini")
	default:
		log.Fatal("\"env\"'s value is not defined")
	}
	if err != nil {
		log.Fatal("Load settings error:", err)
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
	DBUser = file.Section("database").Key("db_user").MustString("root")
	DBPassword = file.Section("database").Key("db_password").MustString("123456")
	DBName = file.Section("database").Key("db_name").MustString("go_news")
}
