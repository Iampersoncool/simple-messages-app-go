package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func ConnectDatabase(dbName, addr, username, password string) (*sql.DB, error) {
	config := mysql.NewConfig()
	config.Addr = addr
	config.DBName = dbName
	config.User = username
	config.Passwd = password

	connector, err := mysql.NewConnector(config)
	if err != nil {
		return nil, err
	}

	return sql.OpenDB(connector), nil
}
