package middleware

import (
	"github.com/bemijonathan/tickets_app/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare( t bool ) gin.HandlerFunc {
	return func (c *gin.Context) {
		if t {
			c.JSON(401, helpers.ServiceResponse{
				Data: "Authorization failed",
				StatusCode: 401,
				Status: false,
			})
		}
	}
}