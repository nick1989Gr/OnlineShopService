package database

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

// DbConn is a pointer to the database. We can use this pointer from various
// packages in order to interact with the database
var DbConn *sql.DB

// SetupDatabase is responsible for opening the database and initializng the pointer
// to the database
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:password123@tcp(127.0.0.1:49154)/OnlineShop")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(3)
	DbConn.SetMaxIdleConns(3)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}