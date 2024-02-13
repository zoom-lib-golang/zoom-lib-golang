// +build integration

package zoom

import (
	"os"
	"testing"
)

func TestWebinarList(t *testing.T) {
	AccountID = os.Getenv("ZOOM_ACCOUNT_ID")
	ClientID  = os.Getenv("ZOOM_CLIENT_ID")
	ClientSecret = os.Getenv("ZOOM_CLIENT_SECRET")

	var (
		primaryUser = os.Getenv("ZOOM_EXAMPLE_EMAIL")
		one         = int(1)
	)

	_, err := ListWebinars(ListWebinarsOptions{
		HostID:     primaryUser,
		PageSize:   &one,
		PageNumber: &one,
	})
	if err != nil {
		t.Fatalf("got error listing webinars: %+v\n", err)
	}
}
