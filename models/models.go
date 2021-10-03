package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique"`
	Age          uint8
	Birthday     time.Time
	ActivatedAt  sql.NullTime
	Registers    []Registered
	Transactions []Transactions
}

type Interests struct {
}

type Events struct {
	gorm.Model
	Name          string
	Price         uint8
	Date          time.Time
	StartTime     time.Time
	EndTime       time.Time
	NumOfDays     uint8
	UserID        int
	User          User
	Registers     []Registered
	RegisterCount int
	Transactions  []Transactions
}

type Registered struct {
	gorm.Model
	Name    string
	EventID int
	UserID  int
}

type Transactions struct {
	gorm.Model
	Reference int
	UserID    int
	EventID   int
	amount    int
}
