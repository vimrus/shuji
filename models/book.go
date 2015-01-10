package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id        int64
	Name      string
	Desc      string
	Type      string
	CreatedBy string
	CreatedAt string
}

func GetUserBooks(account string) []Book {
	var books []Book
	_, err := dbmap.Select(&books, "select * from book where createdby = ?", account)
	if err != nil {
		panic(err)
	}
	return books
}
