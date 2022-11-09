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

	listMeetingsOpt := zoom.ListMeetingsOptions{
		HostID: userID,
	}

	listMeetingsResponse, err := zoom.ListMeetings(listMeetingsOpt)

	if err != nil {
		log.Printf("Error: %+v\n\n", err)
	}
	log.Printf("ListMeetings: %+v\n\n", listMeetingsResponse)
}
