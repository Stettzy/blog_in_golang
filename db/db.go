package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Get() (*sql.DB, error) {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", "root:root@/go")
		if err != nil {
			return nil, err
		}

		db.SetConnMaxLifetime(time.Minute * 30)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	}

	return db, nil
}

func Init() error {
	_, err := Get()
	return err
}
