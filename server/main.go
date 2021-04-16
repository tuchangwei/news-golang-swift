package main

import (
	"server/db"
	"server/router"
	"server/utils/settings"
)

func main() {
	settings.InitSettings()
	db.InitDB()
	router.InitRouter()
}
