package db

import (
	"fmt"
	"github.com/bemijonathan/tickets_app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var Db *gorm.DB

func Connect () {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
			dbHost, username, password, dbName,
			)

	fmt.Println(dsn)

	conn, err := gorm.Open("postgres", "\"postgres://root:root@localhost:5432/tickets_app?sslmode=disable\"")
	if err != nil {
		log.Fatal(err.Error())
	}
	Db = conn
	Db.AutoMigrate(&models.User{}, models.Events{}, models.Transactions{})
}