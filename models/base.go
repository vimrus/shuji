package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	//"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func getDB() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:password@/shuji")
	if err != nil {
		panic(err)
	}

	return &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
}
