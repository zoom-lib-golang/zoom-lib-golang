package zoom

import (
	"fmt"
)

// APIError contains the code and message returned by any Zoom errors
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"errors,omitempty"`
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}

	return fmt.Sprintf("Zoom API error %d: \"%s\"", e.Code, e.Message)
}
