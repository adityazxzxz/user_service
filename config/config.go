package config

import (
	"github.com/adityazxzxz/user_service/models"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB{
	db, err := gorm.Open("mysql","root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTimeTrue&loc=Local")
	if err != nil{
		panic("failed to connect to database")
	}

	db.AutoMigrate(models.Users{})

	return db
}