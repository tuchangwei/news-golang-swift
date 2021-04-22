package db

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"server/utils/settings"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
func setup() {
	fmt.Println("_________test start__________")
	settings.InitSettings()
	InitDB()
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
	DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Post{})
}
func shutdown() {
	fmt.Println("_________test end__________")
}
