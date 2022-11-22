package configs

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "wodnjs12",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "onestep",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}
