package routes

import "net/http"

func Status() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UP"))
	}
}
