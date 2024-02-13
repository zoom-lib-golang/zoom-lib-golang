// +build integration foo

package zoom

import (
	"os"
	"testing"
)

func TestListAllRecordings(t *testing.T) {
	AccountID = os.Getenv("ZOOM_ACCOUNT_ID")
	ClientID  = os.Getenv("ZOOM_CLIENT_ID")
	ClientSecret = os.Getenv("ZOOM_CLIENT_SECRET")

	var (
		primaryUser = os.Getenv("ZOOM_EXAMPLE_EMAIL")
		one         = int(1)
	)

	_, err := ListAllRecordings(ListAllRecordingsOptions{
		UserID:   primaryUser,
		PageSize: &one,
		From:     "2019-11-01",
		To:       "2019-12-01",
	})

	if err != nil {
		t.Fatalf("got error listing recordings: %+v\n", err)
	}
}

func TestGetMeetingRecordingsNoRecording(t *testing.T) {
	AccountID = os.Getenv("ZOOM_ACCOUNT_ID")
	ClientID  = os.Getenv("ZOOM_CLIENT_ID")
	ClientSecret = os.Getenv("ZOOM_CLIENT_SECRET")

	_, err := GetMeetingRecordings(GetMeetingRecordingsOptions{
		MeetingID: "FooBar",
	})

	expected := `Zoom API error 3301: "This recording does not exist."`
	actual := err.Error()

	if actual != expected {
		t.Errorf(`Expected error "%s". Got "%s".`, expected, actual)
	}
}
