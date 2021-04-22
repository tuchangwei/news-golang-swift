package main

import (
	"log"
	"server/db"
	"server/router"
	"server/utils/settings"
)

func main() {
	settings.InitSettings()
	db.InitDB()
	router := router.NewRouter()
	err := router.Engine.Run(":"+ settings.HttpPort)
	if err != nil {
		log.Fatal("can't start server", err)
	}
}
