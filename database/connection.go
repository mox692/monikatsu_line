package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func SetupDB() {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	// connectionName := os.Getenv("INSTANCE_CONNECTION_NAME")
	database := os.Getenv("DB_NAME")
	driverName := os.Getenv("DB_DRIVER")

	var err error

	log.Printf("cloudSQLConnection: %s", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, database))
	Conn, err = sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, database))
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	log.Printf("sql connection success!!")
}
