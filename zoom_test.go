package zoom

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	now    Time
	nowStr = "2006-01-02T15:04:05-07:00"
)

func TestMain(m *testing.M) {
	// set set time value to 2006-01-02T15:04:05-07:00
	t, err := time.Parse(time.RFC3339, nowStr)
	now = Time{t}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing test time %s: %s\n", nowStr, err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}
