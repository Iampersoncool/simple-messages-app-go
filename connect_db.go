package main

import (
	"database/sql"
)

func ConnectDatabase(connUrl string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
