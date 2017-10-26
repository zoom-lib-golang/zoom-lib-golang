package zoom

import (
	"fmt"
	"strings"
	"time"
)

// Time is a custom Time type that accounts for null values and empty strings // during JSON marshaling and unmarshaling
type Time struct {
	time.Time
}

// UnmarshalJSON describes JSON unmarshaling for custom Time objects, handling
// empty string values
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		t.Time = time.Time{}
		return
	}

	t.Time, err = time.Parse(time.RFC3339, s)
	return
}

// MarshalJSON describes JSON unmarshaling for custom Time objects, handling
// empty string values
func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(time.RFC3339))), nil
}

// Format calls format on the underlying time object
func (t *Time) Format(format string) string {
	return t.Time.Format(format)
}

// String defines how time is printed out
func (t *Time) String() string {
	return t.Format(time.RFC3339)
}
