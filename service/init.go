package service

import (
	"chatProject/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root:Z00a0319@tcp(139.196.137.117:4306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if nil != err && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DB = db
	DB.AutoMigrate(model.User{}, model.Contact{}, model.Community{})

	fmt.Println("init data base ok")
}
