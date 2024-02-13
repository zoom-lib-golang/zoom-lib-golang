package main

import (
	"log"
	"os"

	"github.com/zoom-lib-golang/zoom-lib-golang"
)

func main() {
	zoom.AccountID = os.Getenv("ZOOM_ACCOUNT_ID")
	zoom.ClientID  = os.Getenv("ZOOM_CLIENT_ID")
	zoom.ClientSecret = os.Getenv("ZOOM_CLIENT_SECRET")

	// Set HostID
	createMeetingOpt := zoom.CreateMeetingOptions{
		HostID: os.Getenv("USER_ID"),
	}

	meeting, err := zoom.CreateMeeting(createMeetingOpt)

	if err != nil {
		log.Fatalf("Error: %+v\n\n", err)
	}
	log.Printf("Meeting: %+v\n\n", meeting)
}
