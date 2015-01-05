package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id        int64
	Name      string
	Password  string
	Email     string
	Avatar    string
	CreatedAt string
	LastLogin string
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

func CreateUser(name string, email string, password string) {
	user := &User{0, name, password, email, "", time.Now().Format("2006-01-02 15:04:05"), ""}
	dbmap.Insert(user)
}
