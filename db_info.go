package main

import "os"

func GetDbInfoFromEnv(defDbName, defUn, defPw string) (string, string, string) {
	dbName := os.Getenv("DB_NAME")
	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	if dbName == "" || userName == "" || password == "" {
		return defDbName, defUn, defPw
	}

	return dbName, userName, password
}
