package main

import (
	"log"
	"os"
	"strconv"

	"github.com/zoom-lib-golang/zoom-lib-golang"
)

func main() {
	var (
		apiKey    = os.Getenv("ZOOM_API_KEY")
		apiSecret = os.Getenv("ZOOM_API_SECRET")
		meetingID = os.Getenv("MEETING_ID")
	)

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret

	id, err := strconv.Atoi(meetingID)

	if err != nil {
		log.Printf("Error: %+v\n\n", err)
	}

	deleteMeetingOpt := zoom.DeleteMeetingOptions{
		MeetingID: id,
	}

	if err = zoom.DeleteMeeting(deleteMeetingOpt); err != nil {
		log.Printf("Error: %+v\n\n", err)
	}
}
