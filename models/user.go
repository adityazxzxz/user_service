package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	First_name string
	Last_name string
}