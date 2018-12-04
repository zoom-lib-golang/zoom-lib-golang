// +build integration

package zoom

import (
	"os"
	"testing"
)

func TestWebinarList(t *testing.T) {
	var (
		apiKey      = os.Getenv("ZOOM_API_KEY")
		apiSecret   = os.Getenv("ZOOM_API_SECRET")
		primaryUser = os.Getenv("ZOOM_EXAMPLE_EMAIL")
		one         = int(1)
	)

	APIKey = apiKey
	APISecret = apiSecret

	_, err := ListWebinars(ListWebinarsOptions{
		HostID:     primaryUser,
		PageSize:   &one,
		PageNumber: &one,
	})
	if err != nil {
		t.Fatalf("got error listing webinars: %+v\n", err)
	}
}
