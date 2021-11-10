package model

import (
	"database/sql"
	_"github.com/lib/pq"
	"fmt"
)

var Db *sql.DB

const (
	driverName string = "postgres"
	dbuser string = "postgres"
	dbName string = "layla"
	dbpassword string = "1234"
	sslMode string = "disable"
)

func init() {
	var err error
	Db, err = sql.Open(driverName, "postgres://postgres:1234@localhost/postgres?sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}












