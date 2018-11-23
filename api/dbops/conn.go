package dbops

import (
	"database/sql"
	"log"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/video_server")
	if err != nil {
		panic(err.Error())
	}
}
