package validators

import (
	"encoding/json"
	"net/http"
)

type ValidationError struct {
	Errors map[string]string `json:"errors"`
}

func NewValidationError(errs map[string]string) *ValidationError {
	return &ValidationError{Errors: errs}
}

func (ve *ValidationError) Error() string {
	errStr, err := json.Marshal(ve)

	if err != nil {
		panic("validationError: failed to json.Marshal error string")
	}

	return string(errStr)
}

func (err *ValidationError) WriteTo(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}
