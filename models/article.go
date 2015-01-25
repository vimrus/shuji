package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Article struct {
	Id        int64
	Title     string
	Content   string
	Order     int
	CreatedBy int64
	CreatedAt string
	LastEdit  string
}

func CreateArticle(title string, userid int64) (*Article, error) {
	article := &Article{0, title, "", 0, userid, time.Now().Format("2006-01-02 15:04:05"), ""}
	err := dbmap.Insert(article)
	return article, err
}
