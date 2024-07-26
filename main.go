package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Iampersoncool/simple-messages-app/routes"
)

func main() {
	db, err := ConnectDatabase(
		GetDbInfoFromEnv("root:password123@tcp(localhost:3306)/simple_messages_app"),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	port := GetPort()
	mux := http.NewServeMux()

	mux.Handle("/status", routes.Status())

	mux.Handle("GET /messages", routes.ListMessages(db))
	mux.Handle("POST /messages", routes.NewMessage(db))

	log.Fatal(http.ListenAndServe(port, mux))
}
