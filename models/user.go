package models

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vimrus/shuji/utils"
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

func GetUserByName(name string) User {
	var user User
	err := dbmap.SelectOne(&user, "select * from user where name = ?", name)
	if err != nil {
		panic(err)
	}
	return user
}

func CreateUser(name string, email string, password string) (*User, error) {
	user := &User{0, name, utils.Sha1(password + name), email, "", time.Now().Format("2006-01-02 15:04:05"), ""}
	err := dbmap.Insert(user)
	return user, err
}

func Authorize(name string, password string) (User, error) {
	var user User
	err := dbmap.SelectOne(&user, "select * from user where name = ?", name)
	if err != nil {
		return user, err
	}
	if user.Password == utils.Sha1(password+user.Name) {
		return user, nil
	}
	return user, errors.New("用户名或者密码错误")
}
