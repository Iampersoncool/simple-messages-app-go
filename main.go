package main

import (
	"log"
	"net/http"

	"github.com/Iampersoncool/simple-messages-app/routes"
)

func main() {
	db, err := ConnectDatabase(
		GetDbInfoFromEnv("localhost", "simple_messages_app", "root", "password123"),
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
