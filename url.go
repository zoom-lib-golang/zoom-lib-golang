package zoom

import (
	"fmt"
	"net/url"
	"strings"
)

// URL is a custom URL type that enables JSON marshaling/unmarshaling of url.URL
type URL struct {
	*url.URL
}

// UnmarshalJSON describes JSON unmarshaling for custom Time objects, handling
// empty string values
func (u *URL) UnmarshalJSON(b []byte) (err error) {
	var ur *url.URL

	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		u.URL = &url.URL{}
		return
	}

	ur, err = url.Parse(s)
	if err != nil {
		return
	}

	u.URL = ur
	return
}

// MarshalJSON describes JSON unmarshaling for custom Time objects, handling
// empty string values
func (u *URL) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", u.String())), nil
}

// String defines how time is printed out
func (u *URL) String() string {
	return u.URL.String()
}
