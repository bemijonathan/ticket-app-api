package db

import (
	"github.com/bemijonathan/tickets_app/models"
	"fmt"
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

	conn, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	Db = conn
	Db.AutoMigrate(&models.User{}, models.Events{}, models.Transactions{})
}