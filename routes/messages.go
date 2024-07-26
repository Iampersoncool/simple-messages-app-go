package routes

import (
	"database/sql"
	"html"
	"log"
	"net/http"

	"github.com/Iampersoncool/simple-messages-app/types"
	"github.com/Iampersoncool/simple-messages-app/utils"
	"github.com/Iampersoncool/simple-messages-app/validators"
)

func NewMessage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := &validators.MessageValidator{
			Message: types.Message{
				Content: r.PostFormValue("content"),
				Color:   r.PostFormValue("color"),
			},
		}

		errs := msg.Validate()
		if errs != nil {
			errs.WriteTo(w)
			return
		}

		stmt, err := db.Prepare("INSERT INTO messages(content, color) VALUES(?, ?)")
		if err != nil {
			log.Printf("Error preparing statement: %v", err)
			utils.WriteJson(w, http.StatusInternalServerError, map[string][]string{
				"errors": {"Error preparing statement"},
			})
			return
		}

		_, execErr := stmt.Exec(
			html.EscapeString(msg.Content),
			html.EscapeString(msg.Color),
		)
		if execErr != nil {
			log.Printf("Error executing statement: %v", err)
			utils.WriteJson(w, http.StatusInternalServerError, map[string][]string{
				"errors": {"Error executing statement"},
			})
			return
		}

		utils.WriteJson(w, http.StatusCreated, map[string]string{
			"message": "Success!",
		})
	}
}

func ListMessages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var msgs []types.Message

		rows, err := db.Query("SELECT content, color FROM messages;")

		if err != nil {
			log.Printf("listMessages route error: %s", err)
			utils.WriteJson(w, http.StatusInternalServerError, map[string][]string{
				"errors": {"Error querying messages"},
			})
			return
		}

		for rows.Next() {
			var msg types.Message
			err := rows.Scan(&msg.Content, &msg.Color)

			if err != nil {
				log.Println(err)

				utils.WriteJson(w, http.StatusInternalServerError, map[string][]string{
					"errors": {"Error querying rows"},
				})

				return
			}

			msgs = append(msgs, msg)
		}

		utils.WriteJson(w, http.StatusOK, map[string][]types.Message{
			"messages": msgs,
		})
	}
}
