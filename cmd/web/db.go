package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	maxOpenDbCon  = 25
	maxIdleDbCon  = 25
	maxDbLifetime = 5 * time.Minute
)

func initMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// testing db connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbCon)
	db.SetMaxIdleConns(maxIdleDbCon)
	db.SetConnMaxLifetime(maxDbLifetime)

	return db, nil
}
