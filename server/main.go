package main

import (
	"server/db"
	"server/router"
)

func main() {
	db.InitDB()
	router.InitRouter()
}
