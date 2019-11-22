// +build integration

package zoom

import (
	"os"
	"testing"
)

func TestListMeetings(t *testing.T) {
	var (
		apiKey      = os.Getenv("ZOOM_API_KEY")
		apiSecret   = os.Getenv("ZOOM_API_SECRET")
		primaryUser = os.Getenv("ZOOM_EXAMPLE_EMAIL")
		one         = int(1)
	)

	APIKey = apiKey
	APISecret = apiSecret

	_, err := ListMeetings(ListMeetingsOptions{
		HostID:     primaryUser,
		PageSize:   &one,
		PageNumber: &one,
	})

	if err != nil {
		t.Fatalf("got error listing meetings: %+v\n", err)
	}
}

func TestCreateGetDeleteMeeting(t *testing.T) {
	var (
		apiKey      = os.Getenv("ZOOM_API_KEY")
		apiSecret   = os.Getenv("ZOOM_API_SECRET")
		primaryUser = os.Getenv("ZOOM_EXAMPLE_EMAIL")
	)

	APIKey = apiKey
	APISecret = apiSecret

	user, err := GetUser(GetUserOpts{EmailOrID: primaryUser})
	if err != nil {
		t.Fatalf("got error listing users: %+v\n", err)
	}

	meeting, err := CreateMeeting(CreateMeetingOptions{
		HostID: user.ID,
		Topic:  "This is a test meeting created by zoom-lib-golang",
		Agenda: "This important topic will be discussed",
		Type:   MeetingTypeInstant,
		Settings: MeetingSettings{
			Audio:            AudioBoth,
			HostVideo:        true,
			ParticipantVideo: false,
			JoinBeforeHost:   true,
			EnforceLogin:     true,
		},
	})
	if err != nil {
		t.Fatalf("got error creating meeting: %+v\n", err)
	}

	meeting, err = GetMeeting(GetMeetingOptions{
		MeetingID: meeting.ID,
	})
	if err != nil {
		t.Fatalf("got error getting meeting: %+v\n", err)
	}

	err = DeleteMeeting(DeleteMeetingOptions{
		MeetingID: meeting.ID,
	})
	if err != nil {
		t.Fatalf("got error getting meeting: %+v\n", err)
	}
}

func TestDeleteMeetingFail(t *testing.T) {
	var (
		apiKey    = os.Getenv("ZOOM_API_KEY")
		apiSecret = os.Getenv("ZOOM_API_SECRET")
	)

	APIKey = apiKey
	APISecret = apiSecret

	err := DeleteMeeting(DeleteMeetingOptions{
		MeetingID: 1234,
	})
	if err == nil {
		t.Fatalf("did not get error getting meeting: %+v\n", err)
	}

	if err.Error() != "404 Not Found" {
		t.Errorf("Expected 404 Not Found. Actual: %v\n", err)
	}
}
