package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var dbmap *gorp.DbMap = getDB()

func getDB() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:password@/shuji")
	if err != nil {
		panic(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(User{}, "user").SetKeys(true, "Id")
	dbmap.AddTableWithName(Book{}, "book").SetKeys(true, "Id")
	return dbmap
}
