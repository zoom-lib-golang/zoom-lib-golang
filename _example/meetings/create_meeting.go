package main

import (
	"log"
	"os"

	"github.com/zoom-lib-golang/zoom-lib-golang"
)

func main() {
	var (
		apiKey    = os.Getenv("ZOOM_API_KEY")
		apiSecret = os.Getenv("ZOOM_API_SECRET")
		userID    = os.Getenv("USER_ID")
	)

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret

	// Set HostID
	createMeetingOpt := zoom.CreateMeetingOptions{
		HostID: userID,
	}

	meeting, err := zoom.CreateMeeting(createMeetingOpt)

	if err != nil {
		log.Printf("Error: %+v\n\n", err)
	}
	log.Printf("Meeting: %+v\n\n", meeting)
}
