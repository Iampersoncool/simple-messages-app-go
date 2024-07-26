package validators

import (
	"strings"

	"github.com/Iampersoncool/simple-messages-app/types"
)

const (
	MaxContentLength = 100
	HexColorLength   = 7
)

type MessageValidator struct {
	types.Message

	Errors map[string]string
}

func (validator *MessageValidator) Validate() *ValidationError {
	validator.Errors = make(map[string]string)

	content := strings.TrimSpace(validator.Content)
	color := strings.TrimSpace(validator.Color)

	if content == "" {
		validator.Errors["content"] = "Please enter content."
	}

	if len(content) > MaxContentLength {
		validator.Errors["content"] = "Content too long."
	}

	if color == "" {
		validator.Errors["color"] = "Please enter a color."
	}

	if len(color) > HexColorLength {
		validator.Errors["content"] = "Not a valid hex color."
	}

	if len(validator.Errors) > 0 {
		return NewValidationError(validator.Errors)
	}

	return nil
}
