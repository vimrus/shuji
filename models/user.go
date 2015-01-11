package models

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vimrus/shuji/utils"
	"time"
)

type User struct {
	Id        int64
	Account   string
	Password  string
	Email     string
	Avatar    string
	CreatedAt string
	LastLogin string
}

func GetUserByAccount(account string) User {
	var user User
	err := dbmap.SelectOne(&user, "select * from user where account = ?", account)
	if err != nil {
		panic(err)
	}
	return user
}

func CreateUser(account string, email string, password string) (*User, error) {
	user := &User{0, account, utils.Sha1(password + account), email, "", time.Now().Format("2006-01-02 15:04:05"), ""}
	err := dbmap.Insert(user)
	return user, err
}

func Authorize(account string, password string) (User, error) {
	var user User
	err := dbmap.SelectOne(&user, "select * from user where account = ?", account)
	if err != nil {
		return user, err
	}
	if user.Password == utils.Sha1(password+user.Account) {
		return user, nil
	}
	return user, errors.New("用户名或者密码错误")
}
