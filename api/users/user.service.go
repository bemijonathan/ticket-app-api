package users

import (
	"github.com/bemijonathan/tickets_app/helpers"
)

func GetUsers () (helpers.ServiceResponse, error)  {
	t := helpers.ServiceResponse{
		Data:       "hello",
		Status:     true,
		StatusCode: 200,
	}
	return t , nil
}