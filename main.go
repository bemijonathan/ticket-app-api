package main

import (
	"github.com/bemijonathan/tickets_app/api/users"
	"github.com/bemijonathan/tickets_app/db"
	middleware_ "github.com/bemijonathan/tickets_app/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func init (){
	//Load .env file
	godotenv.Load()
}

func main(){
	db.Connect()
	router := gin.Default()

	User := router.Group("/user")
	{
		User.GET("/", middleware_.AuthMiddleWare(true), users.GetUserRoute)
	}

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"ping": "pong"})
	})

	router.Run(os.Getenv("port"))

}


