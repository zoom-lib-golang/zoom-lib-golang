package zoom

import "net/http"

// HTTPMethod is the HTTP request method types
type HTTPMethod string

const (
	// Get is GET HTTP method
	Get HTTPMethod = http.MethodGet

	// Post is POST HTTP method
	Post HTTPMethod = http.MethodPost

	// Put is PUT HTTP method
	Put HTTPMethod = http.MethodPut

	// Patch is PATCH HTTP method
	Patch HTTPMethod = http.MethodPatch

	// Delete is DELETE HTTP method
	Delete HTTPMethod = http.MethodDelete
)
