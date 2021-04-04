package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"server/utils"
)
var DB *gorm.DB
func InitDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DBUser, utils.DBPassword, utils.DBHost, utils.DBPort, utils.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("open db error:", err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("auto migrate db error:", err)
	}
	DB = db
}