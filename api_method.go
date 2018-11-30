package zoom

import "net/http"

// Method is the HTTP request method types
type Method string

const (
	// Get is GET HTTP method
	Get Method = http.MethodGet

	// Post is POST HTTP method
	Post Method = http.MethodPost

	// Put is PUT HTTP method
	Put Method = http.MethodPut

	// Patch is PATCH HTTP method
	Patch Method = http.MethodPatch

	// Delete is DELETE HTTP method
	Delete Method = http.MethodDelete
)
