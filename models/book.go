package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Book struct {
	Id        int64
	Name      string
	Desc      string
	Type      string
	CreatedBy int64
	CreatedAt string
}

func CreateBook(name string, desc string, userid int64) (*Book, error) {
	book := &Book{0, name, desc, "", userid, time.Now().Format("2006-01-02 15:04:05")}
	err := dbmap.Insert(book)
	return book, err
}

func GetUserBooks(userid int64) []Book {
	var books []Book
	_, err := dbmap.Select(&books, "select * from book where createdby = ?", userid)
	if err != nil {
		panic(err)
	}
	return books
}
