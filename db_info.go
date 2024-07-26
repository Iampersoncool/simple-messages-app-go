package main

import "os"

func GetDbInfoFromEnv(defAddr, defDbName, defUn, defPw string) (string, string, string, string) {
	addr := os.Getenv("DB_ADDR")
	dbName := os.Getenv("DB_NAME")
	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	if dbName == "" || userName == "" || password == "" || addr == "" {
		return defAddr, defDbName, defUn, defPw
	}

	return addr, dbName, userName, password
}
