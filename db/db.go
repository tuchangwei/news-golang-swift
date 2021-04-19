package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"server/model"
	"server/utils/settings"
)
var DB *gorm.DB
func InitDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.DBUser, settings.DBPassword, settings.DBHost, settings.DBPort, settings.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("open db error:", err)
	}
	err = db.AutoMigrate(&model.User{}, &model.Post{})
	if err != nil {
		log.Fatal("auto migrate db error:", err)
	}
	DB = db
}