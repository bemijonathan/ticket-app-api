package users

import (
	"github.com/bemijonathan/tickets_app/helpers"
	"net/http"
)

func UserRoutes (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response helpers.ServiceResponse
	//var err error

	switch r.Method {
		case "GET":
			response , _ = GetUsers(r)
			break
		case "POST":
	}

	helpers.FormatResponse(w, response)
}