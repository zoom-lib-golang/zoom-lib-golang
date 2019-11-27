// +build integration foo

package zoom

import (
	"os"
	"testing"
)

func TestListAllRecordings(t *testing.T) {
	var (
		apiKey      = os.Getenv("ZOOM_API_KEY")
		apiSecret   = os.Getenv("ZOOM_API_SECRET")
		primaryUser = os.Getenv("ZOOM_EXAMPLE_EMAIL")
		one         = int(1)
	)

	APIKey = apiKey
	APISecret = apiSecret

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
	var (
		apiKey    = os.Getenv("ZOOM_API_KEY")
		apiSecret = os.Getenv("ZOOM_API_SECRET")
	)

	APIKey = apiKey
	APISecret = apiSecret

	_, err := GetMeetingRecordings(GetMeetingRecordingsOptions{
		MeetingID: "FooBar",
	})

	expected := `Zoom API error 3301: "There is no recording for this meeting"`
	actual := err.Error()

	if actual != expected {
		t.Errorf(`Expected error "%s". Got "%s".`, expected, actual)
	}
}
