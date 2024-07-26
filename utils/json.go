package utils

import (
	"encoding/json"
	"net/http"
)

type JH[V any] map[string]V

func WriteJson[V any](w http.ResponseWriter, status int, data JH[V]) {
	result, err := json.Marshal(data)
	if err != nil {
		panic("writeJson error")
	}

	w.WriteHeader(status)
	w.Write(result)
}
