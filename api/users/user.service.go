package users

import (
	"github.com/bemijonathan/tickets_app/helpers"
	"net/http"
)

func GetUsers (r *http.Request) (helpers.ServiceResponse, error)  {
	//println(r)
	t := helpers.ServiceResponse{
		Data:       "hello",
		Status:     true,
		StatusCode: 200,
	}
	return t , nil
}