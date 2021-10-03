package main

import (
	"fmt"
	"github.com/bemijonathan/tickets_app/api/users"
	"github.com/bemijonathan/tickets_app/db"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init (){
	//Load .env file
	godotenv.Load()
}

func main(){

	db.Connect()

	http.HandleFunc("/user",  users.UserRoutes)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new server!")
	})

	port := ":" + os.Getenv("port")
	
	log.Printf("server is starting on %s " , port)
	//log any error
	http.ListenAndServe(
		port,
		nil,
	)
	
}


