package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import the sqlite3 driver

)

var DB *sql.DB

func connDB() error {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func closeConn() error {
	return DB.Close()
}

func setupDB() error {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS tasks (id INTEGER NOT NULL PRIMARY KEY, title TEXT, completed BOOLEAN DEFAULT false, position INTEGER);`)
	if err != nil {
		return err
	}
	return nil
}
