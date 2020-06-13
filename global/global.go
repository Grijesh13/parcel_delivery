package global

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	dbUser     = "root"
	dbPassword = "JustDoIt1308!"
	dbHost     = "localhost"
	dbPort     = "3306"
)

// DB is the SQL client
var DB = initSQL()

func initSQL() *sql.DB {
	// open up a database connection
	db, err:= sql.Open("mysql",
		dbUser + ":" + dbPassword + "@(" + dbHost + ":" + dbPort + ")/")
	if err != nil {
		fmt.Println("Connection Failed!!")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
		panic(err)
	}
	// set configs
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Second * 10)
	return db
}
