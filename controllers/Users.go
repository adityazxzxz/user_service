package controllers

import (
	"net/http"
	"../models"
	"github.com/gin-gonic/gin"
)

func(idb *InDB) GetUser(c *gin.Context){
	var(
		user structs.Users
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?",id).First(&user).Error
	if err != nil{
		result = gin.H{
			"result":err.Error(),
			"count":0
		}
	}else{
		result = gin.H{
			"result":person,
			"count":1,
		}
	}

	c.JSON(http.StatusOK,result)
}

func(idb *InDB) GetUsers(c *.gin.Context){
	var (
		users []structs.Users
		result gin.H
	)

	idb.DB.Find(&users)
	if len(users) <= 0 {
		result = gin.H{
			"result":nil,
			"count":0,
		}
	}else{
		result = gin.H{
			"result":users,
			"count":len(users),
		}
	}

	c.JSON(http.StatusOK,result)
}

func (idb *InDB) CreateUser(c *gin.Context){
	var (
		user structs.Users
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostFomr("last_name")
	user.First_name = first_name
	user.Last_name = last_name
	idb.DB.Create(&user)
	result = gin.H{
		"result":user,
	}
	c.JSON(http.StatusOK,result)
}