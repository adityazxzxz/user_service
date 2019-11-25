package controllers

import (
	"net/http"
	"github.com/adityazxzxz/user_service/models"
	"github.com/gin-gonic/gin"
)

func(idb *InDB) GetUser(c *gin.Context){
	var(
		user models.Users
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?",id).First(&user).Error
	if err != nil{
		result = gin.H{
			"result":err.Error(),
			"count":0,
		}
	}else{
		result = gin.H{
			"result":user,
			"count":1,
		}
	}

	c.JSON(http.StatusOK,result)
}

func(idb *InDB) GetUsers(c *gin.Context){
	var (
		users []models.Users
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
		user models.Users
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	user.First_name = first_name
	user.Last_name = last_name
	idb.DB.Create(&user)
	result = gin.H{
		"result":user,
	}
	c.JSON(http.StatusOK,result)
}

func (idb *InDB) UpdateUser(c *gin.Context){
	var (
		user models.Users
		result gin.H
	)
	id := c.Param("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")

	
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result":"data not found"
		}
	}

	user.First_name = first_name
	user.Last_name = last_name
	err = idb.DB.Model(&user).Updates(user).Error

	if err != nil{
		result = gin.H{
			"result" : "update failed"
		}
	}else{
		result = gin.H{
			"result" : "Successfully updated"
		}
	}

	c.JSON(http.StatusOK,result)
	

}