package zoom

import (
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"
)

var (
	now    Time
	nowStr = "2006-01-02T15:04:05-07:00"

	today    Date
	todayStr = "2006-01-02"

	urlT   URL
	urlStr = "https://zoom.us"
)

func TestMain(m *testing.M) {
	// set the time value to 2006-01-02T15:04:05-07:00
	t, err := time.Parse(time.RFC3339, nowStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing test time %s: %s\n", nowStr, err)
		os.Exit(1)
	}
	now = Time{t}

	// set the date value to 2006-01-02
	d, err := time.Parse(DateFormat, todayStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing test date %s: %s\n", todayStr, err)
		os.Exit(1)
	}
	today = Date{d}

	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing test url %s: %s\n", urlStr, err)
		os.Exit(1)
	}
	urlT = URL{u}

	os.Exit(m.Run())
}
