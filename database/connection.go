package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/appengine"
)

var Conn *sql.DB

func SetupDB() {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	connectionName := os.Getenv("INSTANCE_CONNECTION_NAME")
	database := os.Getenv("DB_DATABASE")
	driverName := os.Getenv("DB_DRIVER")

	var err error

	switch appengine.IsAppEngine() {
	case true:
		cloudSQLConnection := user + ":" + pass + "@unix(/cloudsql/" + connectionName + ")/" + database + "?parseTime=true"
		Conn, err = sql.Open("mysql", cloudSQLConnection)
		if err != nil {
			log.Println(err)
			panic(err.Error())
		}
	case false:
		Conn, err = sql.Open(driverName,
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, database))
		if err != nil {
			log.Println(err)
			panic(err.Error())
		}
	}
}
