package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic"
)

var client *elastic.Client = NewClient()
var dbmap *gorp.DbMap = getDB()

func NewClient() *elastic.Client {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	// Check if shuji exists.
	exists, err := client.IndexExists("shuji").Do()
	if err != nil {
		panic(err)
	}

	if !exists {
		_, err := client.CreateIndex("shuji").Do()
		if err != nil {
			panic(err)
		}
	}
	return client
}

func getDB() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:password@/shuji")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(User{}, "user").SetKeys(true, "Id")
	dbmap.AddTableWithName(Book{}, "book").SetKeys(true, "Id")
	return dbmap
}
