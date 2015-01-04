package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id        int64
	Name      string
	Nick      string
	Password  string
	Email     string
	Avatar    string
	CreatedAt *time.Time
	LastLogin *time.Time
}

func GetUserByName(name string) string {
	db := getDB()
	count, err := db.SelectInt("select count(*) from user")
	if err != nil {
		panic(err)
	}
	count = count + 1
	return "i"
}
