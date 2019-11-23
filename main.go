package main
import (
	"github.com/adityazxzxz/user_service/config"
	"github.com/adityazxzxz/user_service/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db := config.DBInit()
	inDB := &controllers.InDB{DB:db}

	router := gin.Default()

	router.GET("/user/:id",InDB.GetUser)
	router.GET("/users",InDB.GetUsers)
	router.POST("/user",InDB.CreateUser)
	router.Run(":3000")
}