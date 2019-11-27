package zoom

import (
	"fmt"
	"strings"
	"time"
)

const (
	// DateFormat is a date only format string
	DateFormat = "2006-01-02"
)

// Date is a custom Date type that accounts for null values and empty strings // during JSON marshaling and unmarshaling
type Date struct {
	time.Time
}

// UnmarshalJSON describes JSON unmarshaling for custom Date objects, handling
// empty string values
func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		d.Time = time.Time{}
		return
	}

	d.Time, err = time.Parse(DateFormat, s)
	return
}

// MarshalJSON describes JSON unmarshaling for custom Date objects, handling
// empty string values
func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format(DateFormat))), nil
}

// Format calls format on the underlying date object
func (d *Date) Format(format string) string {
	return d.Time.Format(format)
}

// String defines how date is printed out
func (d *Date) String() string {
	return d.Format(DateFormat)
}
