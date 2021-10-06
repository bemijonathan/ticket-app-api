package helpers

import (
	"github.com/gin-gonic/gin"
)

func FormatResponse (c * gin.Context ,  data ServiceResponse) {
	c.JSON(data.StatusCode, data )
}
