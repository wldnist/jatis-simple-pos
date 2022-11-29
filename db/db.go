package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	HOST = "localhost"
	PORT = 3306
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		username, password, HOST, PORT, database)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	log.Println("Database connection established")
	return db, nil
}
