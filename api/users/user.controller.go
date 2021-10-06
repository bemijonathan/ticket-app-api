package users

import (
	"github.com/bemijonathan/tickets_app/helpers"
	"github.com/gin-gonic/gin"
)

func GetUserRoute(c *gin.Context) {
			response , _ := GetUsers()
	helpers.FormatResponse(c, response)
}