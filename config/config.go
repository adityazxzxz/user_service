package config

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB{
	db, err := gorm.Open("mysql","root:@tcp(1270.0.0.1:3306)/godb?charset=utf&parseTimeTrue&loc=Local")
	if err != nil{
		panic("failedto connect to database")
	}

	db.AutoMigrate(models.Users{})

	return db
}